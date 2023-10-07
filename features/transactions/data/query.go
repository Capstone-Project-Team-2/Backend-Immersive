package data

import (
	"bytes"
	"capstone-tickets/apps/config"
	_eventData "capstone-tickets/features/events/data"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

var log = helpers.Log()

type transactionQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) transactions.TransactionDataInterface {
	return &transactionQuery{
		db: database,
	}
}

// Insert implements transactions.TransactionDataInterface.
func (r *transactionQuery) Insert(data transactions.TransactionCore, buyer_id string) error {
	//1. memulai transaksi
	tx := r.db.Begin()
	//2. pengecekan stok tiket
	var transactionModel = TransactionCoreToModel(data)
	fmt.Println("data:", data)
	fmt.Println("transactionModel:", transactionModel)
	var count = map[string]int{}
	var paymentTotal float64
	for i, v := range transactionModel.TicketDetail {
		var errGen error
		transactionModel.TicketDetail[i].ID, errGen = helpers.GenerateUUID()
		if errGen != nil {
			return errGen
		}
		fmt.Println(errGen)
		transactionModel.TicketDetail[i].BuyerID = buyer_id
		fmt.Println("ticket:", v.TicketID)
		_, exist := count[v.TicketID]
		if !exist {
			count[v.TicketID] = 1
		} else {
			count[v.TicketID] += 1
		}
	}

	fmt.Println("map :", count)
	for key, v := range count {
		var ticket _eventData.Ticket
		//3.pengecekan waktu penjualan tiket
		err := tx.Where("id =? and sell_start <= NOW() and sell_end > NOW()", key).First(&ticket).Error
		if err != nil {
			tx.Rollback()
			return errors.New("tiket tidak tersedia " + err.Error())
		}
		// if tx1.RowsAffected == 0 {
		// 	return errors.New("tiket tidak tersedia " + tx1.Error.Error())
		// }
		//4. hitung total pembayaran dan stok tiket
		fmt.Println("ticketprice: ", ticket.Price)
		paymentTotal = paymentTotal + (float64(ticket.Price) * float64(v))
		if ticket.Total < uint(v) {
			tx.Rollback()
			return errors.New("tiket tidak mencukupi" + err.Error())
		}
	}
	fmt.Println(paymentTotal)

	uuid, err := helpers.GenerateUUID()
	if err != nil {
		tx.Rollback()
		return err
	}
	transactionModel.ID = uuid
	transactionModel.PaymentTotal = paymentTotal + 5000.00
	transactionModel.BuyerID = buyer_id

	var bank = BankTransfer{
		Bank: "bca",
	}
	var trans = TransactionDetail{
		OrderID:     transactionModel.ID,
		GrossAmount: transactionModel.PaymentTotal,
	}
	var midtrans = DataMidtrans{
		PaymentType:       "bank_transfer",
		TransactionDetail: trans,
		BankTransfer:      bank,
	}

	jsonData, errMars := json.Marshal(midtrans)
	if errMars != nil {
		tx.Rollback()
		return errMars
	}

	//5. kirim ke midtrans
	request, errReq := http.NewRequest("POST", "https://api.sandbox.midtrans.com/v2/charge", bytes.NewBuffer(jsonData))
	if errReq != nil {
		tx.Rollback()
		return errReq
	}

	request.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1PM1p5QWt2bmgteHByQ0czS0p1OG1OM2w=")
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, errResp := client.Do(request)
	if errResp != nil {
		tx.Rollback()
		return errResp
	}

	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		tx.Rollback()
		return errRead
	}
	fmt.Println("amount", transactionModel.PaymentTotal)
	fmt.Println("id", transactionModel.ID)
	fmt.Println("ada pesan:")
	fmt.Println(string(body))

	var midtransresp MidtransResponse
	//7. ambil response midtrans
	json.Unmarshal(body, &midtransresp)

	transactionModel.VirtualAccount = midtransresp.VirtualAccount[0].VANumber
	transactionModel.TimeLimit = helpers.ParseTimeMidtrans(midtransresp.ExpiredTime)

	for key, v := range count {
		var ticket _eventData.Ticket
		err := tx.Where("id=?", key).First(&ticket).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		//8. pengurangan ticket
		err = tx.Model(&_eventData.Ticket{}).Where("id = ?", key).Update("total", ticket.Total-uint(v)).Error
		fmt.Println("total:", ticket.Total, v, key, ticket.Total-uint(v), count)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//9. simpan ke database
	err = tx.Create(&transactionModel).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//10. commit transaksi
	// tx.Rollback()
	defer tx.Commit()
	return nil
}

// Select implements transactions.TransactionDataInterface.
func (r *transactionQuery) Select(id string) (transactions.TransactionCore, error) {
	var transaction Transaction

	tx := r.db.Where("id=?", id).First(&transaction)
	if tx.Error != nil {
		return transactions.TransactionCore{}, tx.Error
	}
	data := TransactionModelToCore(transaction)
	return data, nil
}

// Update implements transactions.TransactionDataInterface.
func (r *transactionQuery) Update(input transactions.MidtransCallbackCore) error {
	tx := r.db.Begin()
	var err error

	var sign = CheckSignatureKey(input.SignatureKey, input.OrderID, input.StatusCode, input.GrossAmount, config.KEY_SERVER)
	if !sign {
		fmt.Println("sign:", sign)
		fmt.Println("midtrans sign:", input.SignatureKey)
		return errors.New("signature unvalid")
	}

	var transaksiUpdate Transaction

	transaksiUpdate.OrderID = input.TransactionID
	transaksiUpdate.PaymentStatus = CheckStatus(input.TransactionStatus, input.FraudStatus)

	if transaksiUpdate.PaymentStatus == "Failed" {
		var trans Transaction
		var count map[string]int
		err = tx.Where("id=?", input.OrderID).Preload("TicketDetail").First(&trans).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, v := range trans.TicketDetail {
			_, check := count[v.ID]
			if !check {
				count[v.ID] = 1
			} else {
				count[v.ID] += 1
			}
		}
		for key, v := range count {
			var ticket _eventData.Ticket
			err = tx.Where("id = ?", key).First(&ticket).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = tx.Model(&_eventData.Ticket{}).Where("id = ?", key).Update("total", ticket.Total+uint(v)).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Model(&Transaction{}).Where("id=?", input.OrderID).Updates(&transaksiUpdate).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	defer tx.Commit()
	return nil
}

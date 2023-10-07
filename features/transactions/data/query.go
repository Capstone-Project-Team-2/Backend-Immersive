package data

import (
	"bytes"
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
		tx1 := r.db.Where("id =? and sell_start <= NOW() and sell_end > NOW()", key).First(&ticket)
		if tx1.Error != nil {
			return errors.New("tiket tidak tersedia " + tx1.Error.Error())
		}
		if tx1.RowsAffected == 0 {
			return errors.New("tiket tidak tersedia " + tx1.Error.Error())
		}
		//4. hitung total pembayaran dan stok tiket
		fmt.Println("ticketprice: ", ticket.Price)
		paymentTotal = paymentTotal + (float64(ticket.Price) * float64(v))
		if ticket.Total < uint(v) {
			return errors.New("tiket tidak mencukupi" + tx1.Error.Error())
		}
	}
	fmt.Println(paymentTotal)

	//5. Membuat transaksi
	var errGen error
	transactionModel.ID, errGen = helpers.GenerateUUID()
	if errGen != nil {
		return errGen
	}
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
		return errMars
	}

	//7. kirim ke midtrans
	request, errReq := http.NewRequest("POST", "https://api.sandbox.midtrans.com/v2/charge", bytes.NewBuffer(jsonData))
	if errReq != nil {
		return errReq
	}

	request.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1PM1p5QWt2bmgteHByQ0czS0p1OG1OM2w=")
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, errResp := client.Do(request)
	if errResp != nil {
		return errResp
	}

	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return errRead
	}
	fmt.Println("amount", transactionModel.PaymentTotal)
	fmt.Println("id", transactionModel.ID)
	fmt.Println("ada pesan:")
	fmt.Println(string(body))

	var midtransresp MidtransResponse

	json.Unmarshal(body, &midtransresp)

	transactionModel.VirtualAccount = midtransresp.VirtualAccount[0].VANumber
	transactionModel.TimeLimit = helpers.ParseTimeMidtrans(midtransresp.ExpiredTime)

	tx := r.db.Begin()
	for key, v := range count {
		var ticket _eventData.Ticket
		tx.Where("id=?", key).First(&ticket)
		if tx.Error != nil {
			// tx.Rollback()
			return tx.Error
		}

		tx.Model(&_eventData.Ticket{}).Where("id = ?", key).Update("total", ticket.Total-uint(v))
		fmt.Println("total:", ticket.Total, v, key, ticket.Total-uint(v), count)
		if tx.Error != nil {
			// tx.Rollback()
			return tx.Error
		}
	}
	//6. simpan ke database
	tx.Create(&transactionModel)
	if tx.Error != nil {
		// tx.Rollback()
		return tx.Error
	}
	tx.Rollback()
	tx.Commit()
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
func (*transactionQuery) Update(id string, updatedData transactions.TransactionCore) error {
	panic("unimplemented")
}

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
		tx1 := r.db.Where("id =?", key).First(&ticket)
		fmt.Println("ticketprice: ", ticket.Price)
		paymentTotal = paymentTotal + (float64(ticket.Price) * float64(v))
		if ticket.Total < uint(v) {
			return errors.New("tiket tidak mencukupi" + tx1.Error.Error())
		}
	}
	fmt.Println(paymentTotal)
	//3.pengecekan waktu event
	var event _eventData.Event
	tx2 := r.db.Where("id = ? AND end_date > NOW()", transactionModel.EventID).First(&event)
	if tx2.RowsAffected == 0 {
		return errors.New("Waktu event sudah berakhir")
	}

	//4. hitung total pembayaran
	// sudah dihandle diatas

	//5. Membuat transaksi
	tx := r.db.Begin()
	var err error
	transactionModel.ID, err = helpers.GenerateUUID()
	if err != nil {
		tx.Rollback()
		return err
	}
	transactionModel.PaymentTotal = paymentTotal + 5000.00
	transactionModel.BuyerID = buyer_id
	//6. simpan ke database
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

	//7. kirim ke midtrans
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

	json.Unmarshal(body, &midtransresp)

	var errParse error
	transactionModel.VirtualAccount = midtransresp.VirtualAccount[0].VANumber
	transactionModel.TimeLimit, errParse = helpers.ParseTime(midtransresp.ExpiredTime)
	if errParse != nil {
		tx.Rollback()
		return errParse
	}

	tx.Create(&transactionModel)
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	// dataResp := TransactionModelToCore(transactionData)
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

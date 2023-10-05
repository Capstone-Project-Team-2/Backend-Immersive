package data

import (
	"bytes"
	_eventData "capstone-tickets/features/events/data"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/midtrans/midtrans-go"
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
func (r *transactionQuery) Insert(data transactions.TransactionCore) (transactions.TransactionCore, error) {

	//1. memulai transaksi

	//2. pengecekan stok tiket
	var transactionModel = TransactionCoreToModel(data)
	var count map[string]int
	var paymentTotal float64
	for _, v := range transactionModel.TicketDetail {
		_, exist := count[v.TicketID]
		if !exist {
			count[v.TicketID] = 1
		} else {
			count[v.TicketID] += 1
		}
	}

	for key, v := range count {
		var ticket _eventData.Ticket
		tx1 := r.db.Where("id =?", key).First(&ticket)
		paymentTotal = paymentTotal + (float64(ticket.Price) * float64(v))
		if ticket.Total < uint(v) {
			return transactions.TransactionCore{}, errors.New("tiket tidak mencukupi" + tx1.Error.Error())
		}
	}

	//3.pengecekan waktu event
	var event _eventData.Event
	tx2 := r.db.Where("id = ? AND end_date < NOW()", transactionModel.EventID).First(&event)
	if tx2.RowsAffected == 0 {
		return transactions.TransactionCore{}, errors.New("Waktu event sudah berakhir")
	}

	//4. hitung total pembayaran
	// sudah dihandle diatas

	//5. Membuat transaksi
	tx := r.db.Begin()
	var err error
	transactionModel.ID, err = helpers.GenerateUUID()
	if err != nil {
		return transactions.TransactionCore{}, err
	}
	transactionModel.PaymentTotal = paymentTotal

	//6. simpan ke database
	tx.Create(&transactionModel)
	if tx.Error != nil {
		tx.Rollback()
		return transactions.TransactionCore{}, tx.Error
	}

	var bank = BankTransfer{
		Bank: "bca",
	}
	var trans = TransactionDetail{
		OrderID:     transactionModel.ID,
		GrossAmount: paymentTotal,
	}
	var midtrans = DataMidtrans{
		PaymentType:       "bank_transfer",
		TransactionDetail: trans,
		BankTransfer:      bank,
	}

	jsonData, errMars := json.Marshal(midtrans)
	if errMars != nil {
		return transactions.TransactionCore{}, errMars
	}

	//7. kirim ke midtrans
	http.Post("https://api.sandbox.midtrans.com/v2/charge", "application/json", bytes.NewBuffer(jsonData))
	_, err = SendTransactionToMidtrans(transactionData)
	if err != nil {
		tx.Rollback()
		return transactions.TransactionCore{}, err
	}

	tx.Commit()
	dataResp := TransactionModelToCore(transactionData)
	return dataResp, nil
}

// SendTransactionToMidtrans sends transaction data to Midtrans
func SendTransactionToMidtrans(transaction Transaction.TransactionCore, paymentMethod string) (*midtrans.TransactionResponse, error) {
	coreGateway := midtrans.CoreGateway{
		ClientKey: midtrans.ClientKey,
		ServerKey: midtrans.ServerKey,
		Env:       midtrans.Sandbox, // or midtrans.Production for production environment
	}

	chargeReq := &midtrans.ChargeReq{
		PaymentType: midtrans.PaymentType(paymentMethod),
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderID,
			GrossAmt: int64(transaction.PaymentTotal),
		},
		// ... konfigurasi lainnya sesuai dengan jenis pembayaran yang Anda gunakan
	}

	// Kirim data transaksi ke Midtrans
	chargeResp, err := coreGateway.ChargeTransaction(chargeReq)
	if err != nil {
		return nil, err
	}

	return chargeResp, nil
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

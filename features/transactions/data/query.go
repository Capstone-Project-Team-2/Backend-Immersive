package data

import (
	"capstone-tickets/features/tickets"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"errors"

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
	tx := r.db.Begin()

	//2. pengecekan stok tiket
	var availableTickets uint
	tx.Model(&Tickets{}).
		Where("event_id=? AND total >= ?", tickets.TicketCore.EventID, tickets.TicketCore.Total).Count(&availableTickets)
	if availableTickets < data.TicketCount {
		tx.Rollback()
		return transactions.TransactionCore{}, errors.New("stok tiket tidak cukup")
	}

	//3.pengecekan waktu event
	var event tickets.TicketCore
	tx.Where("id = ? AND end_date < NOW()", tickets.TicketCore.EventID)
	if event.ID == "" {
		tx.Rollback()
		return transactions.TransactionCore{}, errors.New("Waktu event sudah berakhir")
	}

	//4. hitung total pembayaran
	ticketPrice := tickets.TicketCore.Price
	//gimana cara membaca perjenis ticketnya
	paymentTotal := float64(data.TicketCount) * ticketPrice

	//5. Membuat transaksi
	transactionData := TransactionCoreToModel(data)
	var err error
	transactionData.ID, err = helpers.GenerateUUID()
	if err != nil {
		return transactions.TransactionCore{}, err
	}
	transactionData.PaymentTotal = paymentTotal

	//6. simpan ke database
	tx = r.db.Create(&transactionData)
	if tx.Error != nil {
		tx.Rollback()
		return transactions.TransactionCore{}, tx.Error
	}

	//7. kirim ke midtrans
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

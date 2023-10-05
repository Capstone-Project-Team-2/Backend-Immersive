package data

import (
	"capstone-tickets/features/tickets"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"errors"

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
	data.PaymentTotal

	transactionData := TransactionCoreToModel(data)
	var err error
	transactionData.ID, err = helpers.GenerateUUID()
	if err != nil {
		return transactions.TransactionCore{}, err
	}

	tx := r.db.Create(&transactionData)
	if tx.Error != nil {
		return transactions.TransactionCore{}, tx.Error
	}
	dataResp := TransactionModelToCore(transactionData)
	return dataResp, nil
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

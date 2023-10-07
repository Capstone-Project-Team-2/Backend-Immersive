package data

import (
	buyerModel "capstone-tickets/features/buyers/data"
	"capstone-tickets/features/transactions"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             string           `gorm:"column:id;type:varchar(191);primaryKey"`
	OrderID        string           `gorm:"column:order_id;type:varchar(191)"`
	BuyerID        string           `gorm:"column:buyer_id;type:varchar(191)"`
	EventID        string           `gorm:"column:event_id;type:varchar(191)"`
	PaymentStatus  string           `gorm:"column:payment_status;type:enum('Pending','Paid', 'Failed');default:Pending"`
	PaymentMethod  string           `gorm:"column:payment_method;"`
	VirtualAccount string           `gorm:"column:virtual_account"`
	TimeLimit      time.Time        `gorm:"column:time_limit"`
	TicketCount    uint             `gorm:"column:ticket_count"`
	PaymentTotal   float64          `gorm:"column:payment_total"`
	Buyer          buyerModel.Buyer `gorm:"foreignKey:BuyerID"`
	TicketDetail   []TicketDetail   `gorm:"foreignKey:TransactionID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type TicketDetail struct {
	ID            string `gorm:"type:varchar(191);primaryKey"`
	BuyerID       string
	EventID       string
	TicketID      string
	TransactionID string
	UseStatus     string `gorm:"column:use_status;type:enum('Pending','Used');default:Pending"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func TransactionModelToCore(transaction Transaction) transactions.TransactionCore {
	return transactions.TransactionCore{
		ID:             transaction.ID,
		OrderID:        transaction.OrderID,
		BuyerID:        transaction.BuyerID,
		EventID:        transaction.EventID,
		PaymentStatus:  transaction.PaymentStatus,
		PaymentMethod:  transaction.PaymentMethod,
		VirtualAccount: transaction.VirtualAccount,
		TimeLimit:      transaction.TimeLimit,
		TicketCount:    transaction.TicketCount,
		PaymentTotal:   transaction.PaymentTotal,
	}
}

func TransactionCoreToModel(core transactions.TransactionCore) Transaction {
	return Transaction{
		ID:             core.ID,
		OrderID:        core.OrderID,
		BuyerID:        core.BuyerID,
		EventID:        core.EventID,
		PaymentStatus:  core.PaymentStatus,
		PaymentMethod:  core.PaymentMethod,
		VirtualAccount: core.VirtualAccount,
		TimeLimit:      core.TimeLimit,
		TicketCount:    core.TicketCount,
		PaymentTotal:   core.PaymentTotal,
		TicketDetail:   TicketDetailCoreToModel(core.TicketDetail),
	}
}

func TicketDetailCoreToModel(input []transactions.TicketDetailCore) []TicketDetail {
	var ticketDetailModel []TicketDetail
	for _, v := range input {
		var ticket = TicketDetail{
			ID:            v.ID,
			BuyerID:       v.BuyerID,
			EventID:       v.EventID,
			TicketID:      v.TicketID,
			TransactionID: v.TransactionID,
			UseStatus:     v.UseStatus,
		}
		ticketDetailModel = append(ticketDetailModel, ticket)
	}
	return ticketDetailModel
}

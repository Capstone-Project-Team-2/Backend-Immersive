package data

import (
	eventModel "capstone-tickets/features/events/data"
	transactionModel "capstone-tickets/features/transactions/data"
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID        string           `gorm:"column:id;type:varchar(191);primaryKey"`
	EventID   string           `gorm:"column:event_id;type:varchar(191)"`
	NameClass string           `gorm:"column:name_class;"`
	Total     uint             `gorm:"column:total"`
	Price     uint             `gorm:"column:price"`
	Event     eventModel.Event `gorm:"foreignKey:EventID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type TicketDetail struct {
	ID            string                       `gorm:"column:id;type:varchar(191);primaryKey"`
	TicketID      string                       `gorm:"column:ticket_id;type:varchar(191)"`
	TransactionID string                       `gorm:"column:transaction_id;type:varchar(191)"`
	UseStatus     string                       `gorm:"column:use_status;type:enum('Pending','Used');default:Pending"`
	Ticket        Ticket                       `gorm:"foreignKey:TicketID"`
	Transaction   transactionModel.Transaction `gorm:"foreignKey:TransactionID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

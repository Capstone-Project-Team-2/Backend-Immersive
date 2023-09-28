package data

import (
	buyerModel "capstone-tickets/features/buyers/data"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            string           `gorm:"column:id;type:varchar(191);primaryKey"`
	BuyerID       string           `gorm:"column:buyer_id;type:varchar(191)"`
	PaymentStatus string           `gorm:"column:payment_status;type:enum('Pending','Paid');default:Pending"`
	PaymentMethod string           `gorm:"column:payment_method;"`
	TimeLimit     time.Time        `gorm:"column:time_limit"`
	TicketCount   uint             `gorm:"column:ticket_count"`
	PaymentTotal  uint             `gorm:"column:payment_total"`
	Buyer         buyerModel.Buyer `gorm:"foreignKey:BuyerID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

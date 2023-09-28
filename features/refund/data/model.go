package data

import (
	transactionModel "capstone-tickets/features/transactions/data"
	"time"

	"gorm.io/gorm"
)

type Refund struct {
	ID                string                       `gorm:"column:id;type:varchar(191);primaryKey"`
	TransactionID     string                       `gorm:"column:transaction_id;type:varchar(191)"`
	RefundDestination string                       `gorm:"column:refund_destination;not null"`
	RefundStatus      string                       `gorm:"column:refund_status;type:enum('Pending','Approved','Declined');default:Pending"`
	PaymentStatus     string                       `gorm:"column:payment_status;type:enum('Pending', 'Paid');default:Pending"`
	PaymentTotal      uint                         `gorm:"column:payment_total"`
	Transaction       transactionModel.Transaction `gorm:"foreignKey:TransactionID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
}

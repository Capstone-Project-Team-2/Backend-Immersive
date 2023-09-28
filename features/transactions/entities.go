package transactions

import "time"

type TransactionCore struct {
	ID            string
	BuyerID       string
	PaymentStatus string
	PaymentMethod string
	TimeLimit     time.Time
	TicketCount   uint
	PaymentTotal  uint
}

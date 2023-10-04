package transactions

import "time"

type TransactionCore struct {
	ID            string
	OrderID       string
	BuyerID       string
	EventID       string
	PaymentStatus string
	PaymentMethod string
	TimeLimit     time.Time
	TicketCount   uint
	PaymentTotal  float64
}
type TransactionDataInterface interface {
	Insert(data TransactionCore) (TransactionCore, error)
	Select(id string) (TransactionCore, error)
	Update(id string, updatedData TransactionCore) error
}
type TransactionServiceInterface interface {
	Create(data TransactionCore) (TransactionCore, error)
	Get(id string) (TransactionCore, error)
	Update(id string, updatedData TransactionCore) error
}

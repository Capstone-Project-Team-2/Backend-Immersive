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
type TransactionDataInterface interface {
	Insert(transactionData TransactionCore) (string, error)
	Select(id string) (TransactionCore, error)
	Update(id string, updatedData TransactionCore) error
}
type TransactionServiceInterface interface {
	Create(transactionData TransactionCore) (string, error)
	Get(id string) (TransactionCore, error)
	Update(id string, updatedData TransactionCore) error
}

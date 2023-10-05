package transactions

import "time"

type TransactionCore struct {
	ID             string
	OrderID        string
	BuyerID        string
	EventID        string
	PaymentStatus  string
	PaymentMethod  string
	VirtualAccount string
	TimeLimit      time.Time
	TicketCount    uint
	PaymentTotal   float64
	TicketDetail   []TicketDetailCore
}

type TicketDetailCore struct {
	ID            string
	BuyerID       string
	EventID       string
	TicketID      string
	TransactionID string
	UseStatus     string
}

type TransactionDataInterface interface {
	Insert(data TransactionCore, buyer_id string) error
	Select(id string) (TransactionCore, error)
	Update(id string, updatedData TransactionCore) error
}
type TransactionServiceInterface interface {
	Create(data TransactionCore, buyer_id string) error
	Get(id string) (TransactionCore, error)
	Update(id string, updatedData TransactionCore) error
}

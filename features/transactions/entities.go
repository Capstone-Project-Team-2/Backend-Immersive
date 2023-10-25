package transactions

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/features/events"
	"time"
)

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
	Buyer          buyers.BuyerCore
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

type MidtransCallbackCore struct {
	TransactionID     string
	TransactionStatus string
	OrderID           string
	FraudStatus       string
	StatusCode        string
	SignatureKey      string
	GrossAmount       string
}

type PaymentMethodCore struct {
	ID         string
	Bank       string
	ServiceFee float64
}

type TransactionDataInterface interface {
	Insert(data TransactionCore, buyer_id string) error
	Select(transaction_id, buyer_id string) (TransactionCore, events.EventCore, error)
	Update(input MidtransCallbackCore) error
	GetAllTicketDetail(buyer_id string) ([]TicketDetailCore, error)
	GetAllPaymentMethod() ([]PaymentMethodCore, error)
	Validation(data TransactionCore) error
	// 1 event valid, 2. waktu penjualan tiket, 3. Kuota ticket
}
type TransactionServiceInterface interface {
	Create(data TransactionCore, buyer_id string) error
	Get(transaction_id, buyer_id string) (TransactionCore, events.EventCore, error)
	Update(input MidtransCallbackCore) error
	GetAllTicketDetail(buyer_id string) ([]TicketDetailCore, error)
	GetAllPaymentMethod() ([]PaymentMethodCore, error)
}

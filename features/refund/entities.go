package refund

type RefundCore struct {
	ID                string
	TransactionID     string
	RefundDestination string
	RefundStatus      string
	PaymentStatus     string
	PaymentTotal      uint
}

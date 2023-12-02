package refund

type RefundCore struct {
	ID                string
	TransactionID     int
	RefundDestination string
	RefundStatus      string
	PaymentStatus     string
	PaymentTotal      float64
}

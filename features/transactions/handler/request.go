package handler

import (
	"capstone-tickets/features/transactions"
	"time"
)

type TransactionRequest struct {
	ID            string    `json:"id" form:"id"`
	OrderID       string    `json:"order_id" form:"order_id"`
	BuyerID       string    `json:"buyer_id" form:"buyer_id"`
	EventID       string    `json:"event_id" form:"event_id"`
	PaymentStatus string    `json:"payment_status" form:"payment_status"`
	PaymentMethod string    `json:"payment_method" form:"payment_method"`
	TimeLimit     time.Time `json:"time_limit" form:"time_limit"`
	TicketCount   uint      `json:"ticket_count" form:"ticket_count"`
	PaymentTotal  float64   `json:"payment_total" form:"payment_total"`
}

func TransactionRequestToCore(input TransactionRequest) transactions.TransactionCore {
	var transactionCore = transactions.TransactionCore{
		ID:            input.ID,
		OrderID:       input.OrderID,
		BuyerID:       input.BuyerID,
		EventID:       input.EventID,
		PaymentStatus: input.PaymentStatus,
		PaymentMethod: input.PaymentMethod,
		TimeLimit:     input.TimeLimit,
		TicketCount:   input.TicketCount,
		PaymentTotal:  input.PaymentTotal,
	}
	return transactionCore
}

package handler

import (
	"capstone-tickets/features/transactions"
	"time"
)

type TransactionRequest struct {
	ID            string         `json:"id" form:"id"`
	OrderID       string         `json:"order_id" form:"order_id"`
	BuyerID       string         `json:"buyer_id" form:"buyer_id"`
	EventID       string         `json:"event_id" form:"event_id"`
	PaymentStatus string         `json:"payment_status" form:"payment_status"`
	PaymentMethod string         `json:"payment_method" form:"payment_method"`
	TimeLimit     time.Time      `json:"time_limit" form:"time_limit"`
	TicketCount   uint           `json:"ticket_count" form:"ticket_count"`
	PaymentTotal  float64        `json:"payment_total" form:"payment_total"`
	TicketDetail  []TicketDetail `json:"ticket_detail" form:"ticket_detail"`
}

type TicketDetail struct {
	EventID  string `json:"event_id" form:"event_id"`
	TicketID string `json:"ticket_id" form:"ticket_id"`
}

type MidtransCallbackRequest struct {
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	FraudStatus       string `json:"fraud_status"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	GrossAmount       string `json:"gross_amount"`
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
		TicketDetail:  TicketDetailRequestToCore(input.TicketDetail),
	}
	return transactionCore
}

func TicketDetailRequestToCore(input []TicketDetail) []transactions.TicketDetailCore {
	var ticketCore []transactions.TicketDetailCore
	for _, val := range input {
		var ticket = transactions.TicketDetailCore{
			EventID:  val.EventID,
			TicketID: val.TicketID,
		}
		ticketCore = append(ticketCore, ticket)
	}
	return ticketCore
}

func MidtransCallbackReqestToCore(input MidtransCallbackRequest) transactions.MidtransCallbackCore {
	var midtrans = transactions.MidtransCallbackCore{
		TransactionID:     input.TransactionID,
		TransactionStatus: input.TransactionStatus,
		OrderID:           input.OrderID,
		FraudStatus:       input.FraudStatus,
		StatusCode:        input.StatusCode,
		SignatureKey:      input.SignatureKey,
		GrossAmount:       input.GrossAmount,
	}
	return midtrans
}

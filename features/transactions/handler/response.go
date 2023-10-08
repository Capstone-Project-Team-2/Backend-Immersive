package handler

import (
	"capstone-tickets/features/transactions"
	"time"
)

type TransactionResponse struct {
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

type TicketDetailresponse struct {
	ID            string `json:"id" form:"id"`
	BuyerID       string `json:"buyer_id" form:"buyer_id"`
	EventID       string `json:"event_id" form:"event_id"`
	TicketID      string `json:"ticket_id" form:"ticket_id"`
	TransactionID string `json:"transaction_id" form:"transaction_id"`
	UseStatus     string `json:"use_status" form:"use_status"`
}

func TransactionCoreToResponse(input transactions.TransactionCore) TransactionResponse {
	var transactionResp = TransactionResponse{
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
	return transactionResp
}

func ListTransactionCoreToResponse(input []transactions.TransactionCore) []TransactionResponse {
	var transactionResp []TransactionResponse
	for _, value := range input {
		var transaction = TransactionResponse{
			ID:            value.ID,
			OrderID:       value.OrderID,
			BuyerID:       value.BuyerID,
			EventID:       value.EventID,
			PaymentStatus: value.PaymentStatus,
			PaymentMethod: value.PaymentMethod,
			TimeLimit:     value.TimeLimit,
			TicketCount:   value.TicketCount,
			PaymentTotal:  value.PaymentTotal,
		}
		transactionResp = append(transactionResp, transaction)
	}
	return transactionResp
}

func ListTicketDetailCoreToResponse(input []transactions.TicketDetailCore) []TicketDetailresponse {
	var ticketDetail []TicketDetailresponse
	for _, v := range input {
		var ticket = TicketDetailresponse{
			ID:            v.ID,
			BuyerID:       v.BuyerID,
			EventID:       v.EventID,
			TicketID:      v.TicketID,
			TransactionID: v.TransactionID,
			UseStatus:     v.UseStatus,
		}
		ticketDetail = append(ticketDetail, ticket)
	}
	return ticketDetail
}

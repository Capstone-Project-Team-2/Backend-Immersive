package handler

import (
	_buyerHandler "capstone-tickets/features/buyers/handler"
	"capstone-tickets/features/transactions"
	"time"
)

type TransactionResponse struct {
	ID             string                      `json:"id,omitempty"`
	OrderID        string                      `json:"order_id,omitempty"`
	BuyerID        string                      `json:"buyer_id,omitempty"`
	EventID        string                      `json:"event_id,omitempty"`
	PaymentStatus  string                      `json:"payment_status,omitempty"`
	PaymentMethod  string                      `json:"payment_method,omitempty"`
	VirtualAccount string                      `json:"virtual_account,omitempty"`
	TimeLimit      time.Time                   `json:"time_limit,omitempty"`
	TicketCount    uint                        `json:"ticket_count,omitempty"`
	PaymentTotal   float64                     `json:"payment_total,omitempty"`
	Buyer          _buyerHandler.BuyerResponse `json:"buyer,omitempty"`
}

type TicketDetailresponse struct {
	ID            string `json:"id,omitempty"`
	BuyerID       string `json:"buyer_id,omitempty"`
	EventID       string `json:"event_id,omitempty"`
	TicketID      string `json:"ticket_id,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	UseStatus     string `json:"use_status,omitempty"`
}

type PaymentMethodResponse struct {
	ID         string  `json:"id,omitempty"`
	Bank       string  `json:"bank,omitempty"`
	ServiceFee float64 `json:"service_fee,omitempty"`
}

func TransactionCoreToResponse(input transactions.TransactionCore) TransactionResponse {
	var transactionResp = TransactionResponse{
		ID:             input.ID,
		OrderID:        input.OrderID,
		BuyerID:        input.BuyerID,
		EventID:        input.EventID,
		PaymentStatus:  input.PaymentStatus,
		PaymentMethod:  input.PaymentMethod,
		VirtualAccount: input.VirtualAccount,
		TimeLimit:      input.TimeLimit,
		TicketCount:    input.TicketCount,
		PaymentTotal:   input.PaymentTotal,
		Buyer:          _buyerHandler.BuyerCoreToResponse(input.Buyer),
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

func ListPaymentMethodCoreToResponse(input []transactions.PaymentMethodCore) []PaymentMethodResponse {
	var paymenMethod []PaymentMethodResponse
	for _, v := range input {
		var payment = PaymentMethodResponse{
			ID:         v.ID,
			Bank:       v.Bank,
			ServiceFee: v.ServiceFee,
		}
		paymenMethod = append(paymenMethod, payment)
	}
	return paymenMethod
}

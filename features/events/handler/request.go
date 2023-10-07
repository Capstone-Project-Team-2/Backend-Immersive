package handler

import (
	"capstone-tickets/features/events"
	"capstone-tickets/helpers"
)

type EventRequest struct {
	Name          string          `json:"name" form:"name" formam:"name"`
	Location      string          `json:"location" form:"location" formam:"location"`
	Description   string          `json:"description" form:"description" formam:"description"`
	TermCondition string          `json:"term_condition" form:"term_condition" formam:"term_condition"`
	StartDate     string          `json:"start_date" form:"start_date" formam:"start_date"`
	EndDate       string          `json:"end_date" form:"end_date" formam:"end_date"`
	Ticket        []TicketRequest `json:"ticket" form:"ticket" formam:"ticket"`
}

type TicketRequest struct {
	ID        string `json:"id" form:"id"`
	NameClass string `json:"name_class" form:"name_class" formam:"name_class"`
	Total     uint   `json:"total" form:"total" formam:"total"`
	Price     uint   `json:"price" form:"price" formam:"price"`
	SellStart string `formam:"sell_start" json:"sell_start" form:"sell_start"`
	SellEnd   string `formam:"sell_end" json:"sell_end" form:"sell_end"`
}

func EventRequestToCore(input EventRequest) events.EventCore {
	var eventCore = events.EventCore{
		Name:          input.Name,
		Location:      input.Location,
		Description:   input.Description,
		TermCondition: input.TermCondition,
		StartDate:     helpers.ParseStringToTime(input.StartDate),
		EndDate:       helpers.ParseStringToTime(input.EndDate),
		Ticket:        ListTicketRequestToCore(input.Ticket),
	}
	return eventCore
}

func ListTicketRequestToCore(input []TicketRequest) []events.TicketCore {
	var ticketsCore []events.TicketCore
	for _, value := range input {
		var ticket = events.TicketCore{
			ID:        value.ID,
			NameClass: value.NameClass,
			Total:     value.Total,
			Price:     value.Price,
			SellStart: helpers.ParseStringToTime(value.SellStart),
			SellEnd:   helpers.ParseStringToTime(value.SellEnd),
		}
		ticketsCore = append(ticketsCore, ticket)
	}
	return ticketsCore
}

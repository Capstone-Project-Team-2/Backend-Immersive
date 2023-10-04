package handler

import (
	"capstone-tickets/features/events"
	"time"
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
	NameClass string `json:"name_class" form:"name_class" formam:"name_class"`
	Total     uint   `json:"total" form:"total" formam:"total"`
	Price     uint   `json:"price" form:"price" formam:"price"`
}

func ParseTime(val string) time.Time {
	layout := "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, val)
	return date
}

func EventRequestToCore(input EventRequest) events.EventCore {
	var eventCore = events.EventCore{
		Name:          input.Name,
		Location:      input.Location,
		Description:   input.Description,
		TermCondition: input.TermCondition,
		StartDate:     ParseTime(input.StartDate),
		EndDate:       ParseTime(input.EndDate),
		Ticket:        ListTicketRequestToCore(input.Ticket),
	}
	return eventCore
}

func ListTicketRequestToCore(input []TicketRequest) []events.TicketCore {
	var ticketsCore []events.TicketCore
	for _, value := range input {
		var ticket = events.TicketCore{
			NameClass: value.NameClass,
			Total:     value.Total,
			Price:     value.Price,
		}
		ticketsCore = append(ticketsCore, ticket)
	}
	return ticketsCore
}

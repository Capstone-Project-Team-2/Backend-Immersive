package handler

import (
	"capstone-tickets/features/events"
	partnerHandler "capstone-tickets/features/partners/handler"
	"time"
)

type EventResponse struct {
	ID            string                         `json:"id" form:"id"`
	Name          string                         `json:"name" form:"name"`
	Location      string                         `json:"location" form:"location"`
	Description   string                         `json:"description" form:"description"`
	TermCondition string                         `json:"term_condition" form:"term_condition"`
	StartDate     time.Time                      `json:"start_date" form:"end_date"`
	EndDate       time.Time                      `json:"end_date" form:"end_date"`
	BannerPicture string                         `json:"banner_picture" form:"banner_picture"`
	Partner       partnerHandler.PartnerResponse `json:"partner" form:"partner"`
	Ticket        []TicketResponse               `json:"ticket" form:"ticket"`
}

type TicketResponse struct {
	ID        string `json:"id" form:"id"`
	EventID   string `json:"event_id" form:"event_id"`
	NameClass string `json:"name_class" form:"name_class"`
	Total     uint   `json:"total" form:"total"`
	Price     uint   `json:"price" form:"price"`
}

func EventCoreToResponse(input events.EventCore) EventResponse {
	var eventResp = EventResponse{
		ID:            input.ID,
		Name:          input.Name,
		Location:      input.Location,
		Description:   input.Description,
		TermCondition: input.TermCondition,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		BannerPicture: input.BannerPicture,
		Partner:       partnerHandler.PartnerCoreToResponse(input.Partner),
		Ticket:        ListTicketCoreToResponse(input.Ticket),
	}
	return eventResp
}

func ListEventCoreToResponse(input []events.EventCore) []EventResponse {
	var eventResp []EventResponse
	for _, value := range input {
		var event = EventResponse{
			ID:            value.ID,
			Name:          value.Name,
			Location:      value.Location,
			Description:   value.Description,
			TermCondition: value.TermCondition,
			StartDate:     value.StartDate,
			EndDate:       value.EndDate,
			BannerPicture: value.BannerPicture,
		}
		eventResp = append(eventResp, event)
	}
	return eventResp
}

func ListTicketCoreToResponse(input []events.TicketCore) []TicketResponse {
	var ticketData []TicketResponse
	for _, value := range input {
		var ticket = TicketResponse{
			ID:        value.ID,
			EventID:   value.EventID,
			NameClass: value.NameClass,
			Total:     value.Total,
			Price:     value.Price,
		}
		ticketData = append(ticketData, ticket)
	}
	return ticketData
}

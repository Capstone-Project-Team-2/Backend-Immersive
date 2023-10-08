package handler

import (
	"capstone-tickets/features/events"
	partnerHandler "capstone-tickets/features/partners/handler"
	"capstone-tickets/helpers"
)

type EventResponse struct {
	ID               string                         `json:"id" form:"id"`
	Name             string                         `json:"name" form:"name"`
	Location         string                         `json:"location" form:"location"`
	Description      string                         `json:"description" form:"description"`
	TermCondition    string                         `json:"term_condition" form:"term_condition"`
	StartDate        string                         `json:"start_date" form:"end_date"`
	EndDate          string                         `json:"end_date" form:"end_date"`
	ValidationStatus string                         `json:"validation_status" form:"validation_status"`
	BannerPicture    string                         `json:"banner_picture" form:"banner_picture"`
	Partner          partnerHandler.PartnerResponse `json:"partner" form:"partner"`
	Ticket           []TicketResponse               `json:"ticket,omitempty" form:"ticket"`
}

type TicketResponse struct {
	ID        string `json:"id" form:"id"`
	EventID   string `json:"event_id" form:"event_id"`
	NameClass string `json:"name_class" form:"name_class"`
	Total     uint   `json:"total" form:"total"`
	Price     uint   `json:"price" form:"price"`
	SellStart string `json:"sell_start"`
	SellEnd   string `json:"sell_end"`
}

func EventCoreToResponse(input events.EventCore) EventResponse {
	var eventResp = EventResponse{
		ID:               input.ID,
		Name:             input.Name,
		Location:         input.Location,
		Description:      input.Description,
		TermCondition:    input.TermCondition,
		StartDate:        helpers.ParseTimeToString(input.StartDate),
		EndDate:          helpers.ParseTimeToString(input.EndDate),
		ValidationStatus: input.ValidationStatus,
		BannerPicture:    helpers.FileFetchEvent + input.BannerPicture,
		Partner:          partnerHandler.PartnerCoreToResponse(input.Partner),
		Ticket:           ListTicketCoreToResponse(input.Ticket),
	}
	return eventResp
}

func ListEventCoreToResponse(input []events.EventCore) []EventResponse {
	var eventResp []EventResponse
	for _, value := range input {
		var event = EventResponse{
			ID:               value.ID,
			Name:             value.Name,
			Location:         value.Location,
			Description:      value.Description,
			TermCondition:    value.TermCondition,
			StartDate:        helpers.ParseTimeToString(value.StartDate),
			EndDate:          helpers.ParseTimeToString(value.EndDate),
			ValidationStatus: value.ValidationStatus,
			BannerPicture:    helpers.FileFetchEvent + value.BannerPicture,
			Partner:          partnerHandler.PartnerCoreToResponse(value.Partner),
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
			SellStart: helpers.ParseTimeToString(value.SellStart),
			SellEnd:   helpers.ParseTimeToString(value.SellEnd),
		}
		ticketData = append(ticketData, ticket)
	}
	return ticketData
}

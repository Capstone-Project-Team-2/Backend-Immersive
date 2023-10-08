package handler

import (
	"capstone-tickets/features/events"
	partnerHandler "capstone-tickets/features/partners/handler"
	"capstone-tickets/helpers"
)

type EventResponse struct {
	ID               string                         `json:"id,omitempty"`
	Name             string                         `json:"name,omitempty"`
	Location         string                         `json:"location,omitempty"`
	Description      string                         `json:"description,omitempty"`
	TermCondition    string                         `json:"term_condition,omitempty"`
	StartDate        string                         `json:"start_date,omitempty"`
	EndDate          string                         `json:"end_date,omitempty"`
	ValidationStatus string                         `json:"validation_status,omitempty"`
	BannerPicture    string                         `json:"banner_picture,omitempty"`
	Partner          partnerHandler.PartnerResponse `json:"partner,omitempty"`
	Ticket           []TicketResponse               `json:"ticket,omitempty"`
}

type TicketResponse struct {
	ID        string `json:"id,omitempty"`
	EventID   string `json:"event_id,omitempty"`
	NameClass string `json:"name_class,omitempty"`
	Total     uint   `json:"total,omitempty"`
	Price     uint   `json:"price,omitempty"`
	SellStart string `json:"sell_start,omitempty"`
	SellEnd   string `json:"sell_end,omitempty"`
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

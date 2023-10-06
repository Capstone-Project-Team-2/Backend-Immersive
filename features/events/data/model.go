package data

import (
	"capstone-tickets/features/events"
	partnerModel "capstone-tickets/features/partners/data"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID               string               `gorm:"column:id;type:varchar(191);primaryKey"`
	PartnerID        string               `gorm:"column:partner_id;type:varchar(191)"`
	Name             string               `gorm:"column:name;not null"`
	Location         string               `gorm:"column:location"`
	Description      string               `gorm:"column:description"`
	TermCondition    string               `gorm:"column:term_condition"`
	StartDate        time.Time            `gorm:"column:start_date"`
	EndDate          time.Time            `gorm:"column:end_date"`
	ValidationStatus string               `gorm:"column:validation_status;type:enum('Pending','Valid','Not Valid');default:Pending"`
	ExecutionStatus  string               `gorm:"column:execution_status;type:enum('On Going','Done');default:'On Going'"`
	BannerPicture    string               `gorm:"column:banner_picture"`
	Partner          partnerModel.Partner `gorm:"foreignKey:PartnerID"`
	Ticket           []Ticket             `gorm:"foreignKey:EventID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type Ticket struct {
	ID        string    `gorm:"column:id;type:varchar(191);primaryKey"`
	EventID   string    `gorm:"column:event_id;type:varchar(191)"`
	NameClass string    `gorm:"column:name_class;"`
	Total     uint      `gorm:"column:total"`
	Price     uint      `gorm:"column:price"`
	SellStart time.Time `gorm:"column:sell_start"`
	SellEnd   time.Time `gorm:"column:sell_end"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func EventCoreToModel(input events.EventCore) Event {
	var eventModel = Event{
		ID:               input.ID,
		PartnerID:        input.PartnerID,
		Name:             input.Name,
		Location:         input.Location,
		Description:      input.Description,
		TermCondition:    input.TermCondition,
		StartDate:        input.StartDate,
		EndDate:          input.EndDate,
		ValidationStatus: input.ValidationStatus,
		ExecutionStatus:  input.ExecutionStatus,
		BannerPicture:    input.BannerPicture,
		Ticket:           ListTicketCoreToModel(input.Ticket),
	}
	return eventModel
}

func EventModelToCore(input Event) events.EventCore {
	var eventCore = events.EventCore{
		ID:               input.ID,
		PartnerID:        input.PartnerID,
		Name:             input.Name,
		Location:         input.Location,
		Description:      input.Description,
		TermCondition:    input.TermCondition,
		StartDate:        input.StartDate,
		EndDate:          input.EndDate,
		ValidationStatus: input.ValidationStatus,
		ExecutionStatus:  input.ExecutionStatus,
		BannerPicture:    input.BannerPicture,
		Partner:          partnerModel.PartnerModelToCore(input.Partner),
		Ticket:           ListTicketModelToCore(input.Ticket),
	}
	return eventCore
}

func ListEventModelToCore(input []Event) []events.EventCore {
	var eventCore []events.EventCore
	for _, value := range input {
		var event = events.EventCore{
			ID:               value.ID,
			PartnerID:        value.PartnerID,
			Name:             value.Name,
			Location:         value.Location,
			Description:      value.Description,
			TermCondition:    value.TermCondition,
			StartDate:        value.StartDate,
			EndDate:          value.EndDate,
			ValidationStatus: value.ValidationStatus,
			ExecutionStatus:  value.ExecutionStatus,
			BannerPicture:    value.BannerPicture,
			Partner:          partnerModel.PartnerModelToCore(value.Partner),
			Ticket:           ListTicketModelToCore(value.Ticket),
		}
		eventCore = append(eventCore, event)
	}
	return eventCore
}

func ListTicketCoreToModel(input []events.TicketCore) []Ticket {
	var ticketModel []Ticket
	for _, value := range input {
		var ticket = Ticket{
			ID:        value.ID,
			NameClass: value.NameClass,
			Total:     value.Total,
			Price:     value.Price,
			SellStart: value.SellStart,
			SellEnd:   value.SellEnd,
		}
		ticketModel = append(ticketModel, ticket)
	}
	return ticketModel
}

func ListTicketModelToCore(input []Ticket) []events.TicketCore {
	var ticketData []events.TicketCore
	for _, value := range input {
		var ticket = events.TicketCore{
			ID:        value.ID,
			EventID:   value.EventID,
			NameClass: value.NameClass,
			Total:     value.Total,
			Price:     value.Price,
			SellStart: value.SellStart,
			SellEnd:   value.SellEnd,
		}
		ticketData = append(ticketData, ticket)
	}
	return ticketData
}

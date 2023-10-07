package events

import (
	"capstone-tickets/features/partners"
	"mime/multipart"
	"time"
)

type EventCore struct {
	ID               string
	PartnerID        string
	Name             string
	Location         string
	Description      string
	TermCondition    string
	StartDate        time.Time
	EndDate          time.Time
	ValidationStatus string
	ExecutionStatus  string
	BannerPicture    string
	Partner          partners.PartnerCore
	Ticket           []TicketCore
}

type TicketCore struct {
	ID        string
	EventID   string
	NameClass string
	Total     uint
	Price     uint
	SellStart time.Time
	SellEnd   time.Time
}

type EventDataInterface interface {
	Insert(input EventCore, file multipart.File) error
	Select(id string) (EventCore, error)
	SelectAll() ([]EventCore, error)
	Update(event_id, partner_id string, input EventCore, file multipart.File) error
	Delete(id string) error
}

type EventServiceInterface interface {
	Add(input EventCore, file multipart.File) error
	Get(id string) (EventCore, error)
	GetAll() ([]EventCore, error)
	Update(event_id, partner_id string, input EventCore, file multipart.File) error
	Delete(id string) error
}

package service

import (
	"capstone-tickets/features/events"
	"mime/multipart"
)

type EventService struct {
	eventData events.EventDataInterface
}

func New(repo events.EventDataInterface) events.EventServiceInterface {
	return &EventService{
		eventData: repo,
	}
}

// Validate implements events.EventServiceInterface.
func (service *EventService) Validate(event_id, validation_status string) error {
	err := service.eventData.Validate(event_id, validation_status)
	return err
}

// Add implements events.EventServiceInterface.
func (service *EventService) Add(input events.EventCore, file multipart.File) error {
	err := service.eventData.Insert(input, file)
	return err
}

// Delete implements events.EventServiceInterface.
func (*EventService) Delete(id string) error {
	panic("unimplemented")
}

// Get implements events.EventServiceInterface.
func (service *EventService) Get(id string) (events.EventCore, error) {
	result, err := service.eventData.Select(id)
	return result, err
}

// GetAll implements events.EventServiceInterface.
func (service *EventService) GetAll(page, item, search string) ([]events.EventCore, bool, error) {
	result, next, err := service.eventData.SelectAll(page, item, search)
	return result, next, err
}

// Update implements events.EventServiceInterface.
func (service *EventService) Update(event_id, partner_id string, input events.EventCore, file multipart.File) error {
	err := service.eventData.Update(event_id, partner_id, input, file)
	return err
}

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
func (*EventService) Get(id string) (events.EventCore, error) {
	panic("unimplemented")
}

// GetAll implements events.EventServiceInterface.
func (service *EventService) GetAll(userId, role, validation, execution string) ([]events.EventCore, error) {
	result, err := service.eventData.SelectAll(userId, role, validation, execution)
	return result, err
}

// Update implements events.EventServiceInterface.
func (*EventService) Update(id string, input events.EventCore) error {
	panic("unimplemented")
}

package data

import (
	"capstone-tickets/features/events"
	"capstone-tickets/helpers"
	"errors"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
)

type EventQuery struct {
	db *gorm.DB
}

var errNoRow = errors.New("no row affected")

func New(db *gorm.DB) events.EventDataInterface {
	return &EventQuery{
		db: db,
	}
}

// Delete implements events.EventDataInterface.
func (*EventQuery) Delete(id string) error {
	panic("unimplemented")
}

// Insert implements events.EventDataInterface.
func (repo *EventQuery) Insert(input events.EventCore, file multipart.File) error {
	var eventModel = EventCoreToModel(input)
	var errGen error
	eventModel.ID, errGen = helpers.GenerateUUID()
	if errGen != nil {
		return errGen
	}

	for i := 0; i < len(eventModel.Ticket); i++ {
		eventModel.Ticket[i].ID, errGen = helpers.GenerateUUID()
		if errGen != nil {
			return errGen
		}
	}

	if eventModel.BannerPicture != helpers.DefaultFile {
		eventModel.BannerPicture = eventModel.ID + eventModel.BannerPicture
		helpers.Uploader.UploadFile(file, eventModel.BannerPicture, helpers.EventPath)
	}
	fmt.Println("query event model: ", eventModel)
	tx := repo.db.Create(&eventModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Select implements events.EventDataInterface.
func (repo *EventQuery) Select(id string) (events.EventCore, error) {
	var eventModel Event
	tx := repo.db.Where("id = ?", id).Preload("Partner").Preload("Ticket").First(&eventModel)
	if tx.Error != nil {
		return events.EventCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return events.EventCore{}, errNoRow
	}
	var eventCore = EventModelToCore(eventModel)
	return eventCore, nil
}

// SelectAll implements events.EventDataInterface.
func (repo *EventQuery) SelectAll() ([]events.EventCore, error) {
	var eventData []Event

	tx := repo.db.Where("execution_status = ? and end_date > NOW()", "On Going").Preload("Partner").Preload("Ticket").Find(&eventData)

	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errNoRow
	}
	var eventCore = ListEventModelToCore(eventData)
	return eventCore, nil
}

// Update implements events.EventDataInterface.
func (*EventQuery) Update(id string, input events.EventCore) error {
	panic("unimplemented")
}

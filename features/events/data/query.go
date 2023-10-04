package data

import (
	"capstone-tickets/features/events"
	"capstone-tickets/helpers"
	"errors"
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

	for _, v := range eventModel.Ticket {
		v.ID, errGen = helpers.GenerateUUID()
		if errGen != nil {
			return errGen
		}
	}

	if eventModel.BannerPicture != helpers.DefaultFile {
		eventModel.BannerPicture = eventModel.ID + eventModel.BannerPicture
		helpers.Uploader.UploadFile(file, eventModel.BannerPicture, helpers.EventPath)
	}

	tx := repo.db.Create(&eventModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Select implements events.EventDataInterface.
func (*EventQuery) Select(id string) (events.EventCore, error) {
	panic("unimplemented")
}

// SelectAll implements events.EventDataInterface.
func (repo *EventQuery) SelectAll(userId, role, validation, execution string) ([]events.EventCore, error) {
	var eventData []Event
	var tx *gorm.DB

	if role == "Buyer" {
		tx = repo.db.Where("execution_status = ?", "On Going").Find(&eventData)
	} else if role == "Partner" {
		tx = repo.db.Where("partner_id = ?", userId).Find(&eventData)
	} else if role == "Admin" {
		tx = repo.db.Find(&eventData)
	}
	// tx = repo.db.Find(&eventData)
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

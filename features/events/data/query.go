package data

import (
	"capstone-tickets/features/events"
	"capstone-tickets/helpers"
	"errors"
	"mime/multipart"
	"strconv"

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

// Validate implements events.EventDataInterface.
func (repo *EventQuery) Validate(event_id, validation_status string) error {
	tx := repo.db.Model(&Event{}).Where("id = ?", event_id).Update("validation_status", validation_status)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errNoRow
	}
	return nil
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
func (repo *EventQuery) SelectAll(page, item, search string) ([]events.EventCore, bool, error) {
	var eventData []Event
	var tx *gorm.DB
	var query = repo.db.Where("execution_status = ? and end_date > NOW() order by created_at desc", "On Going").Preload("Partner")

	if search != "" {
		query = query.Where("name like ?", "%"+search+"%")
	}

	queryCount := query
	tx = queryCount.Find(&eventData)
	if tx.Error != nil {
		return nil, false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, false, errNoRow
	}

	count := tx.RowsAffected
	var pageConv, itemConv int
	var errPage, errItem error
	if page != "" && item != "" {
		pageConv, errPage = strconv.Atoi(page)
		if errPage != nil {
			return nil, false, errPage
		}
		itemConv, errItem = strconv.Atoi(item)
		if errItem != nil {
			return nil, false, errItem
		}
		limit := itemConv
		offset := itemConv * (pageConv - 1)
		query = query.Limit(limit).Offset(offset)
	}
	tx = query.Find(&eventData)
	if tx.Error != nil {
		return nil, false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, false, errNoRow
	}

	next := true
	if itemConv != 0 {
		var totalPage = count / int64(itemConv)
		if count%int64(itemConv) != 0 {
			totalPage += 1
		}
		if totalPage == int64(pageConv) {
			next = false
		}
	}
	var eventCore = ListEventModelToCore(eventData)
	return eventCore, next, nil
}

// Update implements events.EventDataInterface.
func (repo *EventQuery) Update(event_id, partner_id string, input events.EventCore, file multipart.File) error {
	var eventModel Event
	// var tickets = eventModel.Ticket
	txFetch := repo.db.Where("id = ?", event_id).First(&eventModel)
	if txFetch.Error != nil {
		return txFetch.Error
	}
	if txFetch.RowsAffected == 0 {
		return errNoRow
	}

	if eventModel.PartnerID != partner_id {
		return errors.New("Unauthorize")
	}

	var eventUpdate = EventCoreToModel(input)
	var tickets = eventUpdate.Ticket
	for i := 0; i < len(tickets); i++ {
		tickets[i].EventID = event_id
	}

	if eventUpdate.BannerPicture != helpers.DefaultFile {
		eventUpdate.BannerPicture = event_id + eventUpdate.BannerPicture
		helpers.Uploader.UploadFile(file, eventUpdate.BannerPicture, helpers.EventPath)
	}
	tx := repo.db.Model(&Event{}).Where("id = ?", event_id).Updates(&eventUpdate)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errNoRow
	}
	return nil
}

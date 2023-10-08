package data

import (
	eventModel "capstone-tickets/features/events/data"
	"capstone-tickets/features/volunteers"
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	ID        string           `gorm:"column:id;type:varchar(191);primaryKey"`
	EventID   string           `gorm:"column:event_id;type:varchar(191)"`
	Name      string           `gorm:"column:name;not null"`
	Email     string           `gorm:"column:email;not null;unique"`
	Password  string           `gorm:"column:password;not null"`
	Event     eventModel.Event `gorm:"foreignKey:EventID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func VolunteerModelToCore(model Volunteer) volunteers.VolunteerCore {
	return volunteers.VolunteerCore{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		EventID:  model.EventID,
	}
}

func VolunteerCoreToModel(core volunteers.VolunteerCore) Volunteer {
	return Volunteer{
		ID:       core.ID,
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
		EventID:  core.EventID,
	}
}

func ListVolunteerModelToCore(input []Volunteer) []volunteers.VolunteerCore {
	var volunteerCore []volunteers.VolunteerCore
	for _, value := range input {
		var volunteer = volunteers.VolunteerCore{
			ID:       value.ID,
			Name:     value.Name,
			Email:    value.Email,
			Password: value.Password,
			EventID:  value.EventID,
		}
		volunteerCore = append(volunteerCore, volunteer)
	}
	return volunteerCore
}

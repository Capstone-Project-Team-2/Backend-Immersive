package data

import (
	eventModel "capstone-tickets/features/events/data"
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	ID        string           `gorm:"column:id;type:varchar(191);primaryKey"`
	EventID   string           `gorm:"column:event_id;type:varchar(191)"`
	Name      string           `gorm:"column:name;not null"`
	Email     string           `gorm:"column:email;not null"`
	Password  string           `gorm:"column:password;not null"`
	Event     eventModel.Event `gorm:"foreignKey:EventID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

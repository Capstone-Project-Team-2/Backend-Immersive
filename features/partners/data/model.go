package data

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	ID          string    `gorm:"column:id;type:varchar(191);primaryKey"`
	Name        string    `gorm:"column:name"`
	StartJoin   time.Time `gorm:"column:start_join"`
	Email       string    `gorm:"column:email;not null"`
	Password    string    `gorm:"column:password;not null"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Address     string    `gorm:"column:address"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

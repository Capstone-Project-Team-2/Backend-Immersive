package data

import (
	"time"

	"gorm.io/gorm"
)

type Buyer struct {
	ID          string `gorm:"column:id;type:varchar(191);primaryKey"`
	Name        string `gorm:"column:name;not null"`
	PhoneNumber string `gorm:"column:phone_number"`
	Email       string `gorm:"column:email;not null"`
	Password    string `gorm:"column:password;not null"`
	Address     string `gorm:"column:address"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

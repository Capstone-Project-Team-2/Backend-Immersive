package data

import (
	"capstone-tickets/features/admins"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID          string         `gorm:"column:id;type:varchar(191);primaryKey"`
	Name        string         `gorm:"column:name;not null"`
	Email       string         `gorm:"column:email;unique;not nul"`
	Password    string         `gorm:"column:password;not null"`
	PhoneNumber string         `gorm:"column:phone_number;unique"`
	Address     string         `gorm:"column:address"`
	Role        string         `gorm:"column:role;type:enum('Admin','Superadmin')"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;"`
}

func AdminCoretoModel(data admins.AdminCore) Admin {
	return Admin{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		Role:        data.Role,
	}
}

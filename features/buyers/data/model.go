package data

import (
	"capstone-tickets/features/buyers"
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
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func ModelToCore(model Buyer) buyers.BuyerCore {
	return buyers.BuyerCore{
		ID:          model.ID,
		Name:        model.Name,
		PhoneNumber: model.PhoneNumber,
		Email:       model.Email,
		Password:    model.Password,
		Address:     model.Address,
	}
}

func CoreToModel(core buyers.BuyerCore) Buyer {
	return Buyer{
		ID:          core.ID,
		Name:        core.Name,
		PhoneNumber: core.PhoneNumber,
		Email:       core.Email,
		Password:    core.Password,
		Address:     core.Address,
	}
}

package data

import (
	"capstone-tickets/features/partners"
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	ID             string    `gorm:"column:id;type:varchar(191);primaryKey"`
	Name           string    `gorm:"column:name"`
	StartJoin      time.Time `gorm:"column:start_join;autoCreateTime"`
	Email          string    `gorm:"column:email;not null;unique"`
	Password       string    `gorm:"column:password;not null"`
	PhoneNumber    string    `gorm:"column:phone_number;unique"`
	Address        string    `gorm:"column:address"`
	ProfilePicture string    `gorm:"column:profile_picture"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func PartnerCoreToModel(input partners.PartnerCore) Partner {
	var partnerModel = Partner{
		ID:             input.ID,
		Name:           input.Name,
		StartJoin:      input.StartJoin,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Address:        input.Address,
		ProfilePicture: input.ProfilePicture,
	}
	return partnerModel
}

func PartnerModelToCore(input Partner) partners.PartnerCore {
	var partnerCore = partners.PartnerCore{
		ID:             input.ID,
		Name:           input.Name,
		StartJoin:      input.StartJoin,
		Email:          input.Email,
		Password:       input.Password,
		PhoneNumber:    input.PhoneNumber,
		Address:        input.Address,
		ProfilePicture: input.ProfilePicture,
	}
	return partnerCore
}

func ListPartnerModelToCore(input []Partner) []partners.PartnerCore {
	var partnerCore []partners.PartnerCore
	for _, value := range input {
		var partner = partners.PartnerCore{
			ID:             value.ID,
			Name:           value.Name,
			StartJoin:      value.StartJoin,
			Email:          value.Email,
			Password:       value.Password,
			PhoneNumber:    value.PhoneNumber,
			Address:        value.Address,
			ProfilePicture: value.ProfilePicture,
		}
		partnerCore = append(partnerCore, partner)
	}
	return partnerCore
}

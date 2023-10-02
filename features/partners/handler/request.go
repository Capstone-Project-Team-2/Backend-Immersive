package handler

import (
	"capstone-tickets/features/partners"
)

type PartnerLoginrequest struct {
	Email    string `json:"email " form:"email"`
	Password string `json:"password" form:"password"`
}

type PartnerRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

func PartnerRequestToCore(input PartnerRequest) partners.PartnerCore {
	var partnerCore = partners.PartnerCore{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
	}
	return partnerCore
}

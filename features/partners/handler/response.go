package handler

import (
	"capstone-tickets/features/partners"
	"time"
)

type PartnerResponse struct {
	ID             string    `json:"id" form:"id"`
	Name           string    `json:"name" form:"name"`
	StartJoin      time.Time `json:"start_join" form:"start_join"`
	Email          string    `json:"email" form:"email"`
	PhoneNumber    string    `json:"phone_number" form:"phone_number"`
	Address        string    `json:"address" form:"address"`
	ProfilePicture string    `json:"profile_picture" form:"profile_picture"`
}

func PartnerCoreToResponse(input partners.PartnerCore) PartnerResponse {
	var partnerResp = PartnerResponse{
		ID:             input.ID,
		Name:           input.Name,
		StartJoin:      input.StartJoin,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Address:        input.Address,
		ProfilePicture: input.ProfilePicture,
	}
	return partnerResp
}

func ListPartnerCoreToResponse(input []partners.PartnerCore) []PartnerResponse {
	var partnerResp []PartnerResponse
	for _, value := range input {
		var partner = PartnerResponse{
			ID:             value.ID,
			Name:           value.Name,
			StartJoin:      value.StartJoin,
			Email:          value.Email,
			PhoneNumber:    value.PhoneNumber,
			Address:        value.Address,
			ProfilePicture: value.ProfilePicture,
		}
		partnerResp = append(partnerResp, partner)
	}
	return partnerResp
}

package handler

import (
	"capstone-tickets/features/partners"
	"capstone-tickets/helpers"
	"time"
)

type PartnerResponse struct {
	ID             string    `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	StartJoin      time.Time `json:"start_join,omitempty"`
	Email          string    `json:"email,omitempty"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	Address        string    `json:"address,omitempty"`
	ProfilePicture string    `json:"profile_picture,omitempty"`
}

func PartnerCoreToResponse(input partners.PartnerCore) PartnerResponse {
	var partnerResp = PartnerResponse{
		ID:             input.ID,
		Name:           input.Name,
		StartJoin:      input.StartJoin,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Address:        input.Address,
		ProfilePicture: helpers.FileFetchParner + input.ProfilePicture,
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
			ProfilePicture: helpers.FileFetchParner + value.ProfilePicture,
		}
		partnerResp = append(partnerResp, partner)
	}
	return partnerResp
}

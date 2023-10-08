package handler

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
)

type BuyerResponse struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	Address        string `json:"address,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

func BuyerCoreToResponse(input buyers.BuyerCore) BuyerResponse {
	var buyerResp = BuyerResponse{
		ID:             input.ID,
		Name:           input.Name,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Address:        input.Address,
		ProfilePicture: helpers.FileFetchBuyer + input.ProfilePicture,
	}
	return buyerResp
}

func ListBuyerCoreToResponse(input []buyers.BuyerCore) []BuyerResponse {
	var buyerResp []BuyerResponse
	for _, value := range input {
		var buyer = BuyerResponse{
			ID:             value.ID,
			Name:           value.Name,
			Email:          value.Email,
			PhoneNumber:    value.PhoneNumber,
			Address:        value.Address,
			ProfilePicture: helpers.FileFetchBuyer + value.ProfilePicture,
		}
		buyerResp = append(buyerResp, buyer)
	}
	return buyerResp
}

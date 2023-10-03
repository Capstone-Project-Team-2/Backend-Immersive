package handler

import (
	"capstone-tickets/features/buyers"
)

type BuyerResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
}

func BuyerCoreToResponse(input buyers.BuyerCore) BuyerResponse {
	var buyerResp = BuyerResponse{
		ID:             input.ID,
		Name:           input.Name,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		Address:        input.Address,
		ProfilePicture: input.ProfilePicture,
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
			ProfilePicture: value.ProfilePicture,
		}
		buyerResp = append(buyerResp, buyer)
	}
	return buyerResp
}

package handler

import "capstone-tickets/features/buyers"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type BuyerRequest struct {
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	//ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}

func BuyerRequestToCore(input BuyerRequest) buyers.BuyerCore {
	var buyerCore = buyers.BuyerCore{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		//ProfilePicture: input.ProfilePicture,
	}
	return buyerCore
}

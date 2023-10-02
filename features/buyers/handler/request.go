package handler

import "capstone-tickets/features/buyers"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type BuyerRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

func BuyerRequestToCore(input BuyerRequest) buyers.BuyerCore {
	var buyerCore = buyers.BuyerCore{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
	}
	return buyerCore
}

package handler

import "capstone-tickets/features/admins"

type AdminRegister struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Role        string `json:"role"`
}

type LoginAdminRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func AdminRequestToCore(data AdminRegister) admins.AdminCore {
	return admins.AdminCore{
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
		Role:        data.Role,
	}
}

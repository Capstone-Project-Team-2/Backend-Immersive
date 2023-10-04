package handler

import "capstone-tickets/features/admins"

type AdminRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func AdminRequestToCore(data AdminRegister) admins.AdminCore {
	return admins.AdminCore{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Email,
		Role:     data.Role,
	}
}

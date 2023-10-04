package handler

import "capstone-tickets/features/volunteers"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type VolunteerRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	EventID  string `json:"event_id" form:"event_id"`
}

func VolunteerRequestToCore(input VolunteerRequest) volunteers.VolunteerCore {
	var volunteerCore = volunteers.VolunteerCore{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		EventID:  input.EventID,
	}
	return volunteerCore
}

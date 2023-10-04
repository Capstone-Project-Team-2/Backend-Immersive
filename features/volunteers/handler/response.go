package handler

import (
	"capstone-tickets/features/volunteers"
)

type VolunteerResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	EventID string `json:"event_id"`
}

func VolunteerCoreToResponse(input volunteers.VolunteerCore) VolunteerResponse {
	var volunteerResp = VolunteerResponse{
		ID:      input.ID,
		Name:    input.Name,
		Email:   input.Email,
		EventID: input.EventID,
	}
	return volunteerResp
}

func ListVolunteerCoreToResponse(input []volunteers.VolunteerCore) []VolunteerResponse {
	var volunteerResp []VolunteerResponse
	for _, value := range input {
		var volunteer = VolunteerResponse{
			ID:      value.ID,
			Name:    value.Name,
			Email:   value.Email,
			EventID: value.EventID,
		}
		volunteerResp = append(volunteerResp, volunteer)
	}
	return volunteerResp
}

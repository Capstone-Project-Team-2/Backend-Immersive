package service

import (
	"capstone-tickets/features/volunteers"
	"capstone-tickets/helpers"
	"errors"

	"github.com/go-playground/validator"
)

var log = helpers.Log()

type VolunteerService struct {
	volunteerRepo volunteers.VolunteerDataInterface
	validate      *validator.Validate
}

func New(repo volunteers.VolunteerDataInterface) volunteers.VolunteerServiceInterface {
	return &VolunteerService{
		volunteerRepo: repo,
		validate:      validator.New(),
	}
}

// Create implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) Create(input volunteers.VolunteerCore) error {
	err := s.validate.Struct(input)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = s.volunteerRepo.Insert(input)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// Login implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) Login(email string, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", errors.New("Email and Password cannot be empty")
	}
	id, token, err := s.volunteerRepo.Login(email, password)
	return id, token, err
}

// DeleteById implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) DeleteById(id string) error {
	err := s.volunteerRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) GetAll(eventId string) ([]volunteers.VolunteerCore, error) {
	result, err := s.volunteerRepo.SelectAll(eventId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetById implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) GetById(id string) (volunteers.VolunteerCore, error) {
	result, err := s.volunteerRepo.Select(id)
	if err != nil {
		return volunteers.VolunteerCore{}, err
	}
	return result, nil
}

// UpdateById implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) UpdateById(id string, input volunteers.VolunteerCore) error {
	err := s.validate.Struct(input)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = s.volunteerRepo.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}

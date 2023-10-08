package service

import (
	"capstone-tickets/features/volunteers"
	"capstone-tickets/helpers"

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
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	err := s.volunteerRepo.Insert(input)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// Login implements volunteers.VolunteerServiceInterface.
func (s *VolunteerService) Login(email string, password string) (string, string, error) {
	loginInput := volunteers.Login{
		Email:    email,
		Password: password,
	}
	errValidate := s.validate.Struct(loginInput)
	if errValidate != nil {
		return "", "", errValidate
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
func (s *VolunteerService) GetAll(eventId string, param volunteers.QueryParam) (bool, []volunteers.VolunteerCore, error) {
	var totalPage int64
	nextPage := true
	count, dataVolunteer, err := s.volunteerRepo.SelectAll(eventId, param)
	if err != nil {
		return true, nil, err
	}
	if count == 0 {
		nextPage = false
	}
	if param.ExistOtherPage || count != 0 {
		totalPage = count / int64(param.LimitPerPage)
		if count%int64(param.LimitPerPage) != 0 {
			totalPage += 1
		}

		if param.Page == int(totalPage) {
			nextPage = false
		}
		if dataVolunteer == nil {
			nextPage = false
		}
	}
	return nextPage, dataVolunteer, nil
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

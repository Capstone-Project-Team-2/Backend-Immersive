package service

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"errors"

	"github.com/go-playground/validator"
)

var log = helpers.Log()

type BuyerService struct {
	buyerRepo buyers.BuyerDataInterface
	validate  *validator.Validate
}

// DeleteById implements buyers.BuyerServiceInterface.
func (*BuyerService) DeleteById(id string) error {
	panic("unimplemented")
}

// GetAll implements buyers.BuyerServiceInterface.
func (*BuyerService) GetAll() ([]buyers.BuyerCore, error) {
	panic("unimplemented")
}

// GetById implements buyers.BuyerServiceInterface.
func (*BuyerService) GetById(id string) (buyers.BuyerCore, error) {
	panic("unimplemented")
}

// UpdateById implements buyers.BuyerServiceInterface.
func (*BuyerService) UpdateById(id string, input buyers.BuyerCore) error {
	panic("unimplemented")
}

func New(repo buyers.BuyerDataInterface) buyers.BuyerServiceInterface {
	return &BuyerService{
		buyerRepo: repo,
		validate:  validator.New(),
	}
}

// Create implements buyers.BuyerServiceInterface.
func (s *BuyerService) Create(input buyers.BuyerCore) error {
	err := s.validate.Struct(input)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = helpers.ValidatePassword(input.Password)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if input.Name == "" {
		log.Error("fullname is required")
		return errors.New("name is required")
	}
	err = s.buyerRepo.Insert(input)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// Login implements buyers.BuyerServiceInterface.
func (s *BuyerService) Login(email string, password string) (buyers.BuyerCore, string, error) {
	if email == "" || password == "" {
		return buyers.BuyerCore{}, "", errors.New("Email and Password cannot be empty")
	}
	dataLogin, token, err := s.buyerRepo.Login(email, password)
	return dataLogin, token, err
}

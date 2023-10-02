package service

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"errors"
	"mime/multipart"

	"github.com/go-playground/validator"
)

type BuyerService struct {
	buyerRepo buyers.BuyerDataInterface
	validate  *validator.Validate
}

func New(repo buyers.BuyerDataInterface) buyers.BuyerServiceInterface {
	return &BuyerService{
		buyerRepo: repo,
		validate:  validator.New(),
	}
}

// Register implements buyers.BuyerServiceInterface.
func (s *BuyerService) Register(input buyers.BuyerCore, file multipart.File) error {
	err := s.validate.Struct(input)
	if err != nil {
		return err
	}
	err = helpers.ValidatePassword(input.Password)
	if err != nil {
		return err
	}
	err = s.buyerRepo.Register(input, file)
	if err != nil {
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

// Deactive implements buyers.BuyerServiceInterface.
func (*BuyerService) Deactive(id string) error {
	panic("unimplemented")
}

// Edit implements buyers.BuyerServiceInterface.
func (*BuyerService) Edit(id string, input buyers.BuyerCore) error {
	panic("unimplemented")
}

// Profile implements buyers.BuyerServiceInterface.
func (*BuyerService) Profile(id string) (buyers.BuyerCore, error) {
	panic("unimplemented")
}

// ReadAll implements buyers.BuyerServiceInterface.
func (*BuyerService) ReadAll() ([]buyers.BuyerCore, error) {
	panic("unimplemented")
}

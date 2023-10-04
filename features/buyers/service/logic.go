package service

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"errors"
	"mime/multipart"

	"github.com/go-playground/validator"
)

var log = helpers.Log()

type BuyerService struct {
	buyerRepo buyers.BuyerDataInterface
	validate  *validator.Validate
}

// DeleteById implements buyers.BuyerServiceInterface.
func (s *BuyerService) DeleteById(id string) error {
	err := s.buyerRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements buyers.BuyerServiceInterface.
func (s *BuyerService) GetAll() ([]buyers.BuyerCore, error) {
	result, err := s.buyerRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetById implements buyers.BuyerServiceInterface.
func (s *BuyerService) GetById(id string) (buyers.BuyerCore, error) {
	result, err := s.buyerRepo.Select(id)
	if err != nil {
		return buyers.BuyerCore{}, err
	}
	return result, nil
}

// UpdateById implements buyers.BuyerServiceInterface.
func (s *BuyerService) UpdateById(id string, input buyers.BuyerCore, file multipart.File) error {
	err := s.buyerRepo.Update(id, input, file)
	if err != nil {
		return err
	}
	return nil
}

func New(repo buyers.BuyerDataInterface) buyers.BuyerServiceInterface {
	return &BuyerService{
		buyerRepo: repo,
		validate:  validator.New(),
	}
}

// Create implements buyers.BuyerServiceInterface.
func (s *BuyerService) Create(input buyers.BuyerCore, file multipart.File) error {
	err := s.validate.Struct(input)
	if err != nil {
		// log.Error(err.Error())
		return err
	}

	err = s.buyerRepo.Insert(input, file)
	if err != nil {
		// log.Error(err.Error())
		return err
	}
	return nil
}

// Login implements buyers.BuyerServiceInterface.
func (s *BuyerService) Login(email string, password string) (string, string, error) {
	if email == "" || password == "" {
		return "", "", errors.New("Email and Password cannot be empty")
	}
	id, token, err := s.buyerRepo.Login(email, password)
	return id, token, err
}

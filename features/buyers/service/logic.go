package service

import (
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
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
func (s *BuyerService) GetAll(param buyers.QueryParam) (bool, []buyers.BuyerCore, error) {
	var totalPage int64
	nextPage := true
	count, dataBuyer, err := s.buyerRepo.SelectAll(param)
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
		if dataBuyer == nil {
			nextPage = false
		}
	}
	return nextPage, dataBuyer, nil
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
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	err := s.buyerRepo.Insert(input, file)
	if err != nil {
		return err
	}
	return nil
}

// Login implements buyers.BuyerServiceInterface.
func (s *BuyerService) Login(email string, password string) (string, string, error) {
	loginInput := buyers.Login{
		Email:    email,
		Password: password,
	}
	errValidate := s.validate.Struct(loginInput)
	if errValidate != nil {
		return "", "", errValidate
	}
	id, token, err := s.buyerRepo.Login(email, password)
	return id, token, err
}

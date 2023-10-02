package service

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/partners"
	"capstone-tickets/helpers"
	"mime/multipart"
)

type PartnerService struct {
	PartnerData partners.PartnerDataInterface
}

func New(repo partners.PartnerDataInterface) partners.PartnerServiceInterface {
	return &PartnerService{
		PartnerData: repo,
	}
}

// Login implements partners.PartnerServiceInterface.
func (service *PartnerService) Login(email string, password string) (string, string, error) {
	id, err := service.PartnerData.Login(email, password)
	if err != nil {
		return "", "", err
	}
	var token, errToken = middlewares.CreateToken(id, "Partner")
	if errToken != nil {
		return "", "", err
	}
	return id, token, err
}

// Add implements partners.PartnerServiceInterface.
func (service *PartnerService) Add(input partners.PartnerCore, file multipart.File) error {
	var errHass error
	input.Password, errHass = helpers.HassPassword(input.Password)
	if errHass != nil {
		return errHass
	}
	err := service.PartnerData.Insert(input, file)
	return err
}

// Delete implements partners.PartnerServiceInterface.
func (*PartnerService) Delete(id string) error {
	panic("unimplemented")
}

// Get implements partners.PartnerServiceInterface.
func (*PartnerService) Get(id string) (partners.PartnerCore, error) {
	panic("unimplemented")
}

// GetAll implements partners.PartnerServiceInterface.
func (service *PartnerService) GetAll() ([]partners.PartnerCore, error) {
	result, err := service.PartnerData.SelectAll()
	return result, err
}

// Update implements partners.PartnerServiceInterface.
func (*PartnerService) Update(id string, input partners.PartnerCore) error {
	panic("unimplemented")
}

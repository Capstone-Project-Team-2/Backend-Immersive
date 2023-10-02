package service

import (
	"capstone-tickets/features/partners"
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
	id, token, err := service.PartnerData.Login(email, password)
	return id, token, err
}

// Add implements partners.PartnerServiceInterface.
func (service *PartnerService) Add(input partners.PartnerCore, file multipart.File) error {
	err := service.PartnerData.Insert(input, file)
	return err
}

// Delete implements partners.PartnerServiceInterface.
func (service *PartnerService) Delete(id string) error {
	err := service.PartnerData.Delete(id)
	return err
}

// Get implements partners.PartnerServiceInterface.
func (service *PartnerService) Get(id string) (partners.PartnerCore, error) {
	result, err := service.PartnerData.Select(id)
	return result, err
}

// GetAll implements partners.PartnerServiceInterface.
func (service *PartnerService) GetAll() ([]partners.PartnerCore, error) {
	result, err := service.PartnerData.SelectAll()
	return result, err
}

// Update implements partners.PartnerServiceInterface.
func (service *PartnerService) Update(id string, input partners.PartnerCore, file multipart.File) error {
	err := service.PartnerData.Update(id, input, file)
	return err
}

package service

import (
	"capstone-tickets/features/admins"
)

type AdminService struct {
	AdminData admins.AdminDataInterface
}

// Register implements admins.AdminServiceInterface.
func (service *AdminService) Register(data admins.AdminCore) error {
	return service.AdminData.Register(data)

}

func New(repo admins.AdminDataInterface) admins.AdminServiceInterface {
	return &AdminService{
		AdminData: repo,
	}
}

package service

import (
	"capstone-tickets/features/admins"
)

type AdminService struct {
	AdminData admins.AdminDataInterface
}

func New(repo admins.AdminDataInterface) admins.AdminServiceInterface {
	return &AdminService{
		AdminData: repo,
	}
}

// Delete implements admins.AdminServiceInterface.
func (*AdminService) Delete(admin_id string) error {
	panic("unimplemented")
}

// Get implements admins.AdminServiceInterface.
func (*AdminService) Get(admin_id string) (admins.AdminCore, error) {
	panic("unimplemented")
}

// GetAll implements admins.AdminServiceInterface.
func (*AdminService) GetAll() ([]admins.AdminCore, error) {
	panic("unimplemented")
}

// Login implements admins.AdminServiceInterface.
func (service *AdminService) Login(email string, password string) (string, string, error) {
	token, id, err := service.AdminData.Login(email, password)
	return token, id, err
}

// Update implements admins.AdminServiceInterface.
func (*AdminService) Update(admin_id string, input admins.AdminCore) error {
	panic("unimplemented")
}

// Register implements admins.AdminServiceInterface.
func (service *AdminService) Register(data admins.AdminCore) error {
	return service.AdminData.Register(data)

}

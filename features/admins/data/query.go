package data

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/admins"
	"capstone-tickets/helpers"
	"errors"

	"gorm.io/gorm"
)

type AdminData struct {
	db *gorm.DB
}

var errNoRow = errors.New("no row affected")

func New(db *gorm.DB) admins.AdminDataInterface {
	return &AdminData{
		db: db,
	}
}

// Delete implements admins.AdminDataInterface.
func (*AdminData) Delete(admin_id string) error {
	panic("unimplemented")
}

// Get implements admins.AdminDataInterface.
func (*AdminData) Get(admin_id string) (admins.AdminCore, error) {
	panic("unimplemented")
}

// GetAll implements admins.AdminDataInterface.
func (*AdminData) GetAll() ([]admins.AdminCore, error) {
	panic("unimplemented")
}

// Login implements admins.AdminDataInterface.
func (repo *AdminData) Login(email string, password string) (string, string, string, error) {
	var adminModel Admin
	tx := repo.db.Where("email = ?", email).First(&adminModel)
	if tx.Error != nil {
		return "", "", "", tx.Error
	}
	check := helpers.CheckPassword(password, adminModel.Password)
	if !check {
		return "", "", "", errors.New("invalid password")
	}
	token, errToken := middlewares.CreateToken(adminModel.ID, adminModel.Role)
	if errToken != nil {
		return "", "", "", errToken
	}
	return token, adminModel.ID, adminModel.Role, nil
}

// Update implements admins.AdminDataInterface.
func (*AdminData) Update(admin_id string, input admins.AdminCore) error {
	panic("unimplemented")
}

// Register implements admins.AdminDataInterface.
func (repo *AdminData) Register(data admins.AdminCore) error {
	var input = AdminCoretoModel(data)
	var errGen error
	input.ID, errGen = helpers.GenerateUUID()
	if errGen != nil {
		return errGen
	}
	hass, errHass := helpers.HassPassword(input.Password)
	if errHass != nil {
		return errHass
	}
	input.Password = hass
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errNoRow
	}
	return nil
}

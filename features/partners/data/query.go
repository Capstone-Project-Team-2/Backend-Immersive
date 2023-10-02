package data

import (
	"capstone-tickets/features/partners"
	"capstone-tickets/helpers"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type PartnerData struct {
	db *gorm.DB
}

func New(db *gorm.DB) partners.PartnerDataInterface {
	return &PartnerData{
		db: db,
	}
}

// Login implements partners.PartnerDataInterface.
func (repo *PartnerData) Login(email string, password string) (string, error) {
	var partnerData Partner
	tx := repo.db.Where("email = ?").First(&partnerData)
	if tx.Error != nil {
		return "", tx.Error
	}
	check := helpers.CheckPassword(password, partnerData.Password)
	if !check {
		return "", errors.New("password invalid")
	}
	if tx.RowsAffected == 0 {
		return "", errors.New("no row affected")
	}
	return partnerData.ID, nil
}

// Delete implements partners.PartnerDataInterface.
func (*PartnerData) Delete(id string) error {
	panic("unimplemented")
}

// Insert implements partners.PartnerDataInterface.
func (repo *PartnerData) Insert(input partners.PartnerCore, file multipart.File) error {
	var partnerModel = PartnerCoreToModel(input)
	var errGen error
	partnerModel.ID, errGen = helpers.GenerateUUID()
	if errGen != nil {
		return errGen
	}

	if partnerModel.ProfilePicture != helpers.DefaultFile {
		partnerModel.ProfilePicture = partnerModel.ID + partnerModel.ProfilePicture
		helpers.Uploader.UploadFile(file, partnerModel.ProfilePicture)
	}

	tx := repo.db.Create(&partnerModel)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements partners.PartnerDataInterface.
func (*PartnerData) Select(id string) (partners.PartnerCore, error) {
	panic("unimplemented")
}

// SelectAll implements partners.PartnerDataInterface.
func (repo *PartnerData) SelectAll() ([]partners.PartnerCore, error) {
	var partner []Partner
	tx := repo.db.Find(&partner)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no row affected")
	}
	var partnerCore = ListPartnerModelToCore(partner)
	return partnerCore, nil
}

// Update implements partners.PartnerDataInterface.
func (*PartnerData) Update(id string, input partners.PartnerCore) error {
	panic("unimplemented")
}

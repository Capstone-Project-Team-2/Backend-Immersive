package data

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/partners"
	"capstone-tickets/helpers"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"

	"gorm.io/gorm"
)

type PartnerData struct {
	db *gorm.DB
}

var errNoRow = errors.New("no row affected")

func New(db *gorm.DB) partners.PartnerDataInterface {
	return &PartnerData{
		db: db,
	}
}

// Login implements partners.PartnerDataInterface.
func (repo *PartnerData) Login(email string, password string) (string, string, error) {
	var partnerData Partner
	tx := repo.db.Where("email = ?", email).First(&partnerData)
	if tx.Error != nil {
		return "", "", tx.Error
	}
	check := helpers.CheckPassword(password, partnerData.Password)
	if !check {
		return "", "", errors.New("password invalid")
	}
	if tx.RowsAffected == 0 {
		return "", "", errNoRow
	}

	token, errToken := middlewares.CreateToken(partnerData.ID, "Partner")
	if errToken != nil {
		return "", "", errToken
	}
	return partnerData.ID, token, nil
}

// Delete implements partners.PartnerDataInterface.
func (repo *PartnerData) Delete(id string) error {
	tx := repo.db.Where("id = ?", id).Delete(&Partner{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errNoRow
	}
	return nil
}

// Insert implements partners.PartnerDataInterface.
func (repo *PartnerData) Insert(input partners.PartnerCore, file multipart.File) error {
	var partnerModel = PartnerCoreToModel(input)
	var errGen, errHass error
	partnerModel.ID, errGen = helpers.GenerateUUID()
	if errGen != nil {
		return errGen
	}
	partnerModel.Password, errHass = helpers.HassPassword(partnerModel.Password)
	if errHass != nil {
		return errHass
	}

	if partnerModel.ProfilePicture != helpers.DefaultFile {
		partnerModel.ProfilePicture = partnerModel.ID + partnerModel.ProfilePicture
		helpers.Uploader.UploadFile(file, partnerModel.ProfilePicture, helpers.PartnerPath)
	}

	tx := repo.db.Create(&partnerModel)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements partners.PartnerDataInterface.
func (repo *PartnerData) Select(id string) (partners.PartnerCore, error) {
	var partnerModel Partner
	tx := repo.db.Where("id = ?", id).First(&partnerModel)
	if tx.Error != nil {
		return partners.PartnerCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return partners.PartnerCore{}, errNoRow
	}
	var partnerCore = PartnerModelToCore(partnerModel)
	return partnerCore, nil
}

// SelectAll implements partners.PartnerDataInterface.
func (repo *PartnerData) SelectAll(page, item, search string) ([]partners.PartnerCore, bool, error) {
	var partner []Partner
	var tx *gorm.DB
	var query = repo.db

	fmt.Println(search)
	if search != "" {
		query = query.Where("name like ?", "%"+search+"%")
	}

	queryCount := query
	tx = queryCount.Find(&partner)
	if tx.Error != nil {
		return nil, false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, false, errNoRow
	}

	count := tx.RowsAffected

	var pageConv, itemConv int
	var errPage, errItem error
	if page != "" && item != "" {
		pageConv, errPage = strconv.Atoi(page)
		if errPage != nil {
			return nil, false, errPage
		}
		itemConv, errItem = strconv.Atoi(item)
		if errItem != nil {
			return nil, false, errItem
		}
		limit := itemConv
		offset := itemConv * (pageConv - 1)
		query = query.Limit(limit).Offset(offset)
		fmt.Println(limit, offset)
	}

	tx = query.Find(&partner)
	if tx.Error != nil {
		return nil, false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, false, errNoRow
	}

	next := true
	if itemConv != 0 {
		var totalPage = count / int64(itemConv)
		if count%int64(itemConv) != 0 {
			totalPage += 1
		}
		if totalPage == int64(pageConv) {
			next = false
		}
	}

	var partnerCore = ListPartnerModelToCore(partner)
	return partnerCore, next, nil
}

// Update implements partners.PartnerDataInterface.
func (repo *PartnerData) Update(id string, input partners.PartnerCore, file multipart.File) error {
	var partnerFetch Partner
	txFetch := repo.db.Where("id = ?", id).First(&partnerFetch)
	if txFetch.Error != nil {
		return txFetch.Error
	}
	if txFetch.RowsAffected == 0 {
		return errNoRow
	}

	var partnerModel = PartnerCoreToModel(input)
	var errHass error
	if partnerModel.Password != "" {
		partnerModel.Password, errHass = helpers.HassPassword(partnerModel.Password)
		if errHass != nil {
			return errHass
		}
	}

	if partnerModel.ProfilePicture != helpers.DefaultFile {
		partnerModel.ProfilePicture = id + partnerModel.ProfilePicture
		helpers.Uploader.UploadFile(file, partnerModel.ProfilePicture, helpers.PartnerPath)
	} else {
		partnerModel.ProfilePicture = partnerFetch.ProfilePicture
	}

	tx := repo.db.Model(&Partner{}).Where("id = ?", id).Updates(&partnerModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errNoRow
	}
	return nil
}

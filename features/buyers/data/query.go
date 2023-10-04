package data

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

var log = helpers.Log()

type buyerQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) buyers.BuyerDataInterface {
	return &buyerQuery{
		db: database,
	}
}

// Insert implements buyers.BuyerDataInterface.
func (r *buyerQuery) Insert(input buyers.BuyerCore, file multipart.File) error {
	NewData := BuyerCoreToModel(input)

	hashPassword, err := helpers.HassPassword(input.Password)
	if err != nil {
		log.Error("error while hashing password")
		return errors.New("error while hashing password")
	}
	NewData.Password = hashPassword

	NewData.ID, err = helpers.GenerateUUID()
	if err != nil {
		return errors.New("error while generete uuid")
	}

	if NewData.ProfilePicture != helpers.DefaultFile {
		NewData.ProfilePicture = NewData.ID + NewData.ProfilePicture
		helpers.Uploader.UploadFile(file, NewData.ProfilePicture)
	}

	tx := r.db.Create(&NewData)
	if tx.Error != nil {
		log.Error("error insert data")
		return errors.New("error insert data")
	}

	if tx.RowsAffected == 0 {
		log.Warn("no buyer has been created")
		return errors.New("no row affected")
	}
	log.Sugar().Infof("new buyer has been created: %s", NewData.Email)
	return nil
}

// Login implements buyers.BuyerDataInterface.
func (r *buyerQuery) Login(email string, password string) (string, string, error) {
	var dataBuyer Buyer
	tx := r.db.Where("email=?", email).First(&dataBuyer)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return "", "", errors.New("invalid email and password")
	}

	if tx.RowsAffected == 0 {
		return "", "", errors.New("no row affected")
	}

	checkPassword := helpers.CheckPassword(password, dataBuyer.Password)
	if !checkPassword {
		return "", "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken(dataBuyer.ID, "Buyer")
	if err != nil {
		return "", "", errors.New("error while creating jwt token")
	}

	return dataBuyer.ID, token, nil
}

// Update implements buyers.BuyerDataInterface.
func (r *buyerQuery) Update(id string, input buyers.BuyerCore, file multipart.File) error {
	var buyer Buyer
	tx := r.db.Where("id = ?", id).First(&buyer)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	//Mapping Buyer to BuyerCore
	updatedBuyer := BuyerCoreToModel(input)
	if updatedBuyer.Password != "" {
		HassPassword, err := helpers.HassPassword(updatedBuyer.Password)
		if err != nil {
			log.Error("error while hashing password")
			return errors.New("error while hashing password")
		}
		updatedBuyer.Password = HassPassword
	}

	if updatedBuyer.ProfilePicture != helpers.DefaultFile {
		updatedBuyer.ProfilePicture = id + updatedBuyer.ProfilePicture
		helpers.Uploader.UploadFile(file, updatedBuyer.ProfilePicture)
	} else {
		updatedBuyer.ProfilePicture = buyer.ProfilePicture
	}

	tx = r.db.Model(&buyer).Updates(updatedBuyer)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update buyer")
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil

}

// Select implements buyers.BuyerDataInterface.
func (r *buyerQuery) Select(id string) (buyers.BuyerCore, error) {
	var buyer Buyer
	tx := r.db.Where("id = ?", id).First(&buyer)
	if tx.Error != nil {
		return buyers.BuyerCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return buyers.BuyerCore{}, errors.New("data not found")
	}
	//Mapping Buyer to BuyerCore
	coreBuyer := BuyerModelToCore(buyer)
	return coreBuyer, nil
}

// SelectAll implements buyers.BuyerDataInterface.
func (r *buyerQuery) SelectAll() ([]buyers.BuyerCore, error) {
	var buyer []Buyer
	tx := r.db.Find(&buyer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}

	//mapping from buyer -> BuyerCore
	coreBuyerSlice := ListBuyerModelToCore(buyer)
	return coreBuyerSlice, nil
}

// Delete implements buyers.BuyerDataInterface.
func (r *buyerQuery) Delete(id string) error {
	tx := r.db.Where("id = ?", id).Delete(&Buyer{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

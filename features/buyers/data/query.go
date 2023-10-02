package data

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/buyers"
	"capstone-tickets/helpers"
	"errors"

	"gorm.io/gorm"
)

var log = helpers.Log()

type buyerQuery struct {
	db *gorm.DB
}

// Delete implements buyers.BuyerDataInterface.
func (*buyerQuery) Delete(id string) error {
	panic("unimplemented")
}

// Select implements buyers.BuyerDataInterface.
func (*buyerQuery) Select(id string) (buyers.BuyerCore, error) {
	panic("unimplemented")
}

// SelectAll implements buyers.BuyerDataInterface.
func (*buyerQuery) SelectAll() ([]buyers.BuyerCore, error) {
	panic("unimplemented")
}

// Update implements buyers.BuyerDataInterface.
func (*buyerQuery) Update(input buyers.BuyerCore) error {
	panic("unimplemented")
}

func New(database *gorm.DB) buyers.BuyerDataInterface {
	return &buyerQuery{
		db: database,
	}
}

// Insert implements buyers.BuyerDataInterface.
func (r *buyerQuery) Insert(input buyers.BuyerCore) error {
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

	// if NewData.ProfilePicture != helpers.DefaultFile {
	// 	NewData.ProfilePicture = NewData.ID + NewData.ProfilePicture
	// 	helpers.Uploader.UploadFile(file, NewData.ProfilePicture)
	// }

	tx := r.db.Create(&NewData)
	if tx.Error != nil {
		log.Error("error insert data")
		return errors.New("error insert data")
	}

	if tx.RowsAffected == 0 {
		log.Warn("no user has been created")
		return errors.New("no row affected")
	}
	log.Sugar().Infof("new user has been created: %s", NewData.Email)
	return nil
}

// Login implements buyers.BuyerDataInterface.
func (r *buyerQuery) Login(email string, password string) (buyers.BuyerCore, string, error) {
	var dataBuyer Buyer
	tx := r.db.Where("email=?", email).First(&dataBuyer)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return buyers.BuyerCore{}, "", errors.New("invalid email and password")
	}

	if tx.RowsAffected == 0 {
		return buyers.BuyerCore{}, "", errors.New("no row affected")
	}

	checkPassword := helpers.CheckPassword(password, dataBuyer.Password)
	if !checkPassword {
		return buyers.BuyerCore{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken(dataBuyer.ID, "Buyer")
	if err != nil {
		return buyers.BuyerCore{}, "", errors.New("error while creating jwt token")
	}

	data := BuyerModelToCore(dataBuyer)
	return data, token, nil
}

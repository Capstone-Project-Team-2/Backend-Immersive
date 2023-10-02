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

// Register implements buyers.BuyerDataInterface.
func (r *buyerQuery) Register(input buyers.BuyerCore, file multipart.File) error {
	NewData := CoreToModel(input)

	hashPassword, err := helpers.HassPassword(input.Password)
	if err != nil {
		log.Error("error while hashing password")
		return errors.New("error while hashing password")
	}
	NewData.Password = hashPassword

	NewData.ID, err = helpers.GenerateUUID()
	if err != nil {
		log.Error("error while generete uuid")
		return errors.New("error while generete uuid")
	}

	if NewData.Avatar != helpers.DefaultFile {
		NewData.Avatar = NewData.ID + NewData.Avatar
		helpers.Uploader.UploadFile(file, NewData.Avatar)
	}

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
	log.Error("user record not found, invalid email and password")
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return buyers.BuyerCore{}, "", errors.New("invalid email and password")
	}

	if tx.RowsAffected == 0 {
		log.Warn("no user has been created")
		return buyers.BuyerCore{}, "", errors.New("no row affected")
	}

	checkPassword := helpers.CheckPassword(password, dataBuyer.Password)
	log.Warn("password does not match")
	if !checkPassword {
		return buyers.BuyerCore{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken(dataBuyer.ID, "Buyer")
	if err != nil {
		log.Error("error while creating jwt token")
		return buyers.BuyerCore{}, "", errors.New("error while creating jwt token")
	}

	log.Sugar().Infof("user has been logged in: %s", dataBuyer.ID)
	data := ModelToCore(dataBuyer)
	return data, token, nil
}

func New(database *gorm.DB) buyers.BuyerDataInterface {
	return &buyerQuery{
		db: database,
	}
}

// Deactive implements buyers.BuyerDataInterface.
func (*buyerQuery) Deactive(id string) error {
	panic("unimplemented")
}

// Edit implements buyers.BuyerDataInterface.
func (*buyerQuery) Edit(input buyers.BuyerCore) error {
	panic("unimplemented")
}

// Profile implements buyers.BuyerDataInterface.
func (*buyerQuery) Profile(id string) (buyers.BuyerCore, error) {
	panic("unimplemented")
}

// ReadAll implements buyers.BuyerDataInterface.
func (*buyerQuery) ReadAll() ([]buyers.BuyerCore, error) {
	panic("unimplemented")
}

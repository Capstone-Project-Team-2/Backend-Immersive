package data

import (
	"capstone-tickets/apps/middlewares"
	"capstone-tickets/features/volunteers"
	"capstone-tickets/helpers"
	"errors"

	"gorm.io/gorm"
)

var log = helpers.Log()

type volunteerQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) volunteers.VolunteerDataInterface {
	return &volunteerQuery{
		db: database,
	}
}

// Insert implements volunteers.VolunteerDataInterface.
func (r *volunteerQuery) Insert(input volunteers.VolunteerCore) error {
	NewData := VolunteerCoreToModel(input)

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

	tx := r.db.Create(&NewData)
	if tx.Error != nil {
		log.Error("error insert data")
		return errors.New("error insert data")
	}

	if tx.RowsAffected == 0 {
		log.Warn("no volunteer has been created")
		return errors.New("no row affected")
	}
	log.Sugar().Infof("new volunteer has been created: %s", NewData.Email)
	return nil
}

// Login implements volunteers.VolunteerDataInterface.
func (r *volunteerQuery) Login(email string, password string) (volunteers.VolunteerCore, string, error) {
	var dataVolunteer Volunteer
	tx := r.db.Where("email=?", email).First(&dataVolunteer)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return volunteers.VolunteerCore{}, "", errors.New("invalid email and password")
	}

	if tx.RowsAffected == 0 {
		return volunteers.VolunteerCore{}, "", errors.New("no row affected")
	}

	checkPassword := helpers.CheckPassword(password, dataVolunteer.Password)
	if !checkPassword {
		return volunteers.VolunteerCore{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken(dataVolunteer.ID, "Volunteer")
	if err != nil {
		return volunteers.VolunteerCore{}, "", errors.New("error while creating jwt token")
	}

	data := VolunteerModelToCore(dataVolunteer)
	return data, token, nil
}

// Select implements volunteers.VolunteerDataInterface.
func (r *volunteerQuery) Select(id string) (volunteers.VolunteerCore, error) {
	var volunteer Volunteer
	tx := r.db.Where("id = ?", id).First(&volunteer)
	if tx.Error != nil {
		return volunteers.VolunteerCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return volunteers.VolunteerCore{}, errors.New("data not found")
	}
	//Mapping Volunteer to VolunteerCore
	coreVolunteer := VolunteerModelToCore(volunteer)
	return coreVolunteer, nil
}

// SelectAll implements volunteers.VolunteerDataInterface.
func (r *volunteerQuery) SelectAll() ([]volunteers.VolunteerCore, error) {
	var volunteer []Volunteer
	tx := r.db.Find(&volunteer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}

	//mapping from volunteer -> VolunteerCore
	coreVolunteerSlice := ListVolunteerModelToCore(volunteer)
	return coreVolunteerSlice, nil
}

// Update implements volunteers.VolunteerDataInterface.
func (r *volunteerQuery) Update(id string, input volunteers.VolunteerCore) error {
	var volunteer Volunteer
	tx := r.db.Where("id = ?", id).First(&volunteer)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	//Mapping Volunteer to VolunteerCore
	updatedVolunteer := VolunteerCoreToModel(input)
	if updatedVolunteer.Password != "" {
		HassPassword, err := helpers.HassPassword(updatedVolunteer.Password)
		if err != nil {
			log.Error("error while hashing password")
			return errors.New("error while hashing password")
		}
		updatedVolunteer.Password = HassPassword
	}

	tx = r.db.Model(&volunteer).Updates(updatedVolunteer)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update volunteer")
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// Delete implements volunteers.VolunteerDataInterface.
func (r *volunteerQuery) Delete(id string) error {
	tx := r.db.Where("id = ?", id).Delete(&Volunteer{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

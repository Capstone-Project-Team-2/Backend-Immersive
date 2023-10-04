package data

import (
	"capstone-tickets/features/admins"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminData struct {
	db *gorm.DB
}

// Register implements admins.AdminDataInterface.
func (repo *AdminData) Register(data admins.AdminCore) error {
	var input = CoretoModel(data)
	input.ID = uuid.New().String()
	tx := repo.db.Create(&input)
	return tx.Error
}

var errNoRow = errors.New("no row affected")

func New(db *gorm.DB) admins.AdminDataInterface {
	return &AdminData{
		db: db,
	}
}

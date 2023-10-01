package data

import (
	"capstone-tickets/features/buyers"

	"gorm.io/gorm"
)

type buyerQuery struct {
	db *gorm.DB
}

// Delete implements buyers.BuyerDataInterface.
func (*buyerQuery) Delete(id string) error {
	panic("unimplemented")
}

// Insert implements buyers.BuyerDataInterface.
func (*buyerQuery) Insert(input buyers.BuyerCore) error {
	panic("unimplemented")
}

// Login implements buyers.BuyerDataInterface.
func (*buyerQuery) Login(email string, password string) (buyers.BuyerCore, error) {
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

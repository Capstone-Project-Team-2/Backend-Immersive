package data

import (
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"

	"gorm.io/gorm"
)

var log = helpers.Log()

type transactionQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) transactions.TransactionDataInterface {
	return &transactionQuery{
		db: database,
	}
}

// Insert implements transactions.TransactionDataInterface.
func (r *transactionQuery) Insert(data transactions.TransactionCore) (transactions.TransactionCore, error) {
	transactionData := TransactionCoreToModel(data)
	var err error
	transactionData.ID, err = helpers.GenerateUUID()
	if err != nil {
		return transactions.TransactionCore{}, err
	}

	tx := r.db.Create(&transactionData)
	if tx.Error != nil {
		return transactions.TransactionCore{}, tx.Error
	}
	dataResp := TransactionModelToCore(transactionData)
	return dataResp, nil
}

// Select implements transactions.TransactionDataInterface.
func (r *transactionQuery) Select(id string) (transactions.TransactionCore, error) {
	var transaction Transaction

	tx := r.db.Where("id=?", id).First(&transaction)
	if tx.Error != nil {
		return transactions.TransactionCore{}, tx.Error
	}
	data := TransactionModelToCore(transaction)
	return data, nil
}

// Update implements transactions.TransactionDataInterface.
func (*transactionQuery) Update(id string, updatedData transactions.TransactionCore) error {
	panic("unimplemented")
}

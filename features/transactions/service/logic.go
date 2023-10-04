package service

import (
	"capstone-tickets/features/transactions"
)

type TransactionService struct {
	transactionRepo transactions.TransactionDataInterface
}

func New(repo transactions.TransactionDataInterface) transactions.TransactionServiceInterface {
	return &TransactionService{
		transactionRepo: repo,
	}
}

// Create implements transactions.TransactionServiceInterface.
func (s *TransactionService) Create(data transactions.TransactionCore) (transactions.TransactionCore, error) {
	result, err := s.transactionRepo.Insert(data)
	if err != nil {
		return transactions.TransactionCore{}, err
	}
	return result, nil
}

// Get implements transactions.TransactionServiceInterface.
func (s *TransactionService) Get(id string) (transactions.TransactionCore, error) {
	result, err := s.transactionRepo.Select(id)
	if err != nil {
		return transactions.TransactionCore{}, err
	}
	return result, nil
}

// Update implements transactions.TransactionServiceInterface.
func (*TransactionService) Update(id string, updatedData transactions.TransactionCore) error {
	panic("unimplemented")
}

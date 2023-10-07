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
func (s *TransactionService) Create(data transactions.TransactionCore, buyer_id string) error {
	err := s.transactionRepo.Insert(data, buyer_id)
	if err != nil {
		return err
	}
	return nil
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
func (s *TransactionService) Update(input transactions.MidtransCallbackCore) error {
	err := s.transactionRepo.Update(input)
	return err
}

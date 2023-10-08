package service

import (
	"capstone-tickets/features/events"
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

// GetAllPaymentMethode implements transactions.TransactionServiceInterface.
func (s *TransactionService) GetAllPaymentMethod() ([]transactions.PaymentMethodCore, error) {
	result, err := s.transactionRepo.GetAllPaymentMethod()
	return result, err
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
func (s *TransactionService) Get(transaction_id, buyer_id string) (transactions.TransactionCore, events.EventCore, error) {
	resultTrans, resultEvent, err := s.transactionRepo.Select(transaction_id, buyer_id)
	if err != nil {
		return transactions.TransactionCore{}, events.EventCore{}, err
	}
	return resultTrans, resultEvent, nil
}

// Update implements transactions.TransactionServiceInterface.
func (s *TransactionService) Update(input transactions.MidtransCallbackCore) error {
	err := s.transactionRepo.Update(input)
	return err
}

// GetAllTicketDetail implements transactions.TransactionServiceInterface.
func (s *TransactionService) GetAllTicketDetail(buyer_id string) ([]transactions.TicketDetailCore, error) {
	result, err := s.transactionRepo.GetAllTicketDetail(buyer_id)
	return result, err
}

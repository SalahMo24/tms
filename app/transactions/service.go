package transactions

import (
	"errors"
	"tms/app/types"
)

type TransactionService struct {
	transactionRepo TransactionRepository
}

func NewTransactionService(transactionRepo TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *TransactionService) Create(tl TransactionCreate) (string, error) {
	if err := validateTransactionAmount(tl.Amount); err != nil {
		return "", err
	}
	if err := validateTransactionType(tl.TransactionType); err != nil {
		return "", err
	}
	id, err := s.transactionRepo.Create(tl)
	if err != nil {
		return "", err
	}

	return id, nil

}

func validateTransactionAmount(amount float64) error {
	if amount < 0 {
		return errors.New("transaction amount should be a non negative value")
	}
	return nil
}
func validateTransactionType(transactionType types.TransactionType) error {
	if !transactionType.Valid() {
		return errors.New("transaction type is not supported")
	}
	return nil
}

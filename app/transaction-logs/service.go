package transactionlogs

import (
	"errors"
	"tms/app/types"
)

type TransactionLogService struct {
	transactionLogRepo TransactionLogRepository
}

func NewTransactionLogService(transactionLogRepo TransactionLogRepository) *TransactionLogService {
	return &TransactionLogService{
		transactionLogRepo: transactionLogRepo,
	}
}

func (s *TransactionLogService) Create(tl TransactionLogCreate) (string, error) {
	if err := validateTransactionAmount(tl.Amount); err != nil {
		return "", err
	}
	if err := validateTransactionType(tl.TransactionType); err != nil {
		return "", err
	}
	id, err := s.transactionLogRepo.Create(tl)
	if err != nil {
		return "", err
	}

	return id, nil

}
func (s *TransactionLogService) UpdateTransactionLogStatus(status types.Status, id string) (string, error) {

	if err := validateTransactionStatus(status); err != nil {
		return "", err
	}
	id, err := s.transactionLogRepo.UpdateStatus(status, id)
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
func validateTransactionStatus(status types.Status) error {
	if !status.Valid() {
		return errors.New("transaction status is not supported")
	}
	return nil
}

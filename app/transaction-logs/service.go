package transactionlogs

import (
	"tms/app/types"
	"tms/utils/validations"
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
	if err := validations.ValidateTransactionAmount(tl.Amount); err != nil {
		return "", err
	}
	if err := validations.ValidateTransactionType(tl.TransactionType); err != nil {
		return "", err
	}
	id, err := s.transactionLogRepo.Create(tl)
	if err != nil {
		return "", err
	}

	return id, nil

}
func (s *TransactionLogService) UpdateTransactionLogStatus(status types.Status, id string) (string, error) {

	if err := validations.ValidateTransactionStatus(status); err != nil {
		return "", err
	}
	id, err := s.transactionLogRepo.UpdateStatus(status, id)
	if err != nil {
		return "", err
	}

	return id, nil

}

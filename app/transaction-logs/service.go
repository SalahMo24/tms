package transactionlogs

import (
	"tms/app/types"
	"tms/utils/assert"
	"tms/utils/validations"
)

type TransactionLogService struct {
	transactionLogRepo TransactionLogRepository
}

type TransactionLogServiceInterface interface {
	Create(tl TransactionLogCreate) (string, error)
	UpdateTransactionLogStatus(status types.Status, id string) (string, error)
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
	assert.NotNil(id, "id should not be nil")
	assert.Type("", id, "id should be a string")
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
	assert.NotNil(id, "id should not be nil")
	assert.Type("", id, "id should be a string")

	return id, nil

}

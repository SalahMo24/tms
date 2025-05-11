package accountbalance

import (
	"tms/utils/assert"
	"tms/utils/validations"
)

type AccountBalanceService struct {
	accountBalancenRepo AccountBalanceRepository
}

func NewAccountBalanceService(accountBalancenRepo AccountBalanceRepository) *AccountBalanceService {
	return &AccountBalanceService{
		accountBalancenRepo: accountBalancenRepo,
	}
}

func (s *AccountBalanceService) Create(tl AccountBalanceCreate) (string, error) {
	if err := validations.ValidateTransactionAmount(tl.Amount); err != nil {
		return "", err
	}
	if err := validations.ValidateTransactionType(tl.TransactionType); err != nil {
		return "", err
	}
	id, err := s.accountBalancenRepo.Create(tl)
	if err != nil {
		return "", err
	}

	assert.NotNil(id, "id should not be nil")
	assert.Type("", id, "id should be a string")

	return id, nil

}

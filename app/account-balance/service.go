package accountbalance

import (
	"errors"
	"tms/app/types"
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
	if err := validateTransactionAmount(tl.Amount); err != nil {
		return "", err
	}
	if err := validateTransactionType(tl.TransactionType); err != nil {
		return "", err
	}
	id, err := s.accountBalancenRepo.Create(tl)
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

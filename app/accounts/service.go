package accounts

import "tms/utils/assert"

type AccountService struct {
	accRepo AccountRepository
}

func NewUserService(accRepo AccountRepository) *AccountService {
	return &AccountService{
		accRepo: accRepo,
	}
}

func (s *AccountService) Create(userId string) (string, error) {
	id, err := s.accRepo.Create(userId)
	if err != nil {
		return "", err
	}

	assert.NotNil(id, "account id should not be nil")
	assert.Type("", id, "account id should be a string")

	return id, nil

}

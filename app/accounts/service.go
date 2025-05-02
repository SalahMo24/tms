package accounts

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

	return id, nil

}

package users

import (
	"errors"
	"tms/utils/assert"
	"tms/utils/validations"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(req UserCreate) (string, error) {

	if err := validations.ValidateName(req.FirstName, "first name"); err != nil {
		return "", err
	}
	if err := validations.ValidateName(req.LastName, "last name"); err != nil {
		return "", err
	}
	if err := validations.ValidatePhoneNumber(req.PhoneNumber); err != nil {
		return "", err
	}
	if err := validations.ValidateSSN(req.SSN); err != nil {
		return "", err
	}

	exists, _ := s.userRepo.Exists(req.SSN)

	if exists != "" {
		return exists, errors.New("user already exists")
	}

	// Create user
	createdUser, err := s.userRepo.Create(UserCreate{

		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		SSN:         req.SSN,
	})
	if err != nil {
		return "", err
	}

	assert.NotNil(createdUser, "user id should not be nil")
	assert.Type("", createdUser, "user id should be a string")

	return createdUser, nil
}

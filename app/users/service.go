package users

import (
	"errors"
	"unicode"
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

	if err := validateName(req.FirstName, "first name"); err != nil {
		return "", err
	}
	if err := validateName(req.LastName, "last name"); err != nil {
		return "", err
	}
	if err := validatePhoneNumber(req.PhoneNumber); err != nil {
		return "", err
	}
	if err := validateSSN(req.SSN); err != nil {
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

	return createdUser, nil
}

func validateName(name, fieldName string) error {
	// Check length
	if len(name) < 3 || len(name) > 50 {
		return errors.New(fieldName + " must be between 3-50 characters")
	}

	// Check for special characters or numbers
	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return errors.New(fieldName + " must contain only letters and spaces")
		}
	}
	return nil
}

func validatePhoneNumber(phone string) error {
	// Check if starts with valid prefix
	validPrefixes := []string{"010", "011", "012", "014"}
	valid := false
	for _, prefix := range validPrefixes {
		if len(phone) >= len(prefix) && phone[:len(prefix)] == prefix {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New("phone number must start with 010, 011, 012, or 014")
	}

	// Check remaining characters are digits
	if len(phone) < 11 {
		return errors.New("phone number must be at least 11 digits")
	}

	// Check the rest after prefix are digits
	for _, c := range phone {
		if c < '0' || c > '9' {
			return errors.New("phone number must contain only digits")
		}
	}

	return nil
}

func validateSSN(ssn string) error {
	// Check length
	if len(ssn) != 14 {
		return errors.New("SSN must be exactly 14 digits")
	}

	// Check all characters are digits
	for _, c := range ssn {
		if c < '0' || c > '9' {
			return errors.New("SSN must contain only digits")
		}
	}

	return nil
}

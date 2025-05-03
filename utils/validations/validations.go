package validations

import (
	"errors"
	"tms/app/types"
	"unicode"
)

func ValidateTransactionAmount(amount float64) error {
	if amount < 0 {
		return errors.New("transaction amount should be a non negative value")
	}
	return nil
}
func ValidateTransactionType(transactionType types.TransactionType) error {
	if !transactionType.Valid() {
		return errors.New("transaction type is not supported")
	}
	return nil
}

func ValidateTransactionStatus(status types.Status) error {
	if !status.Valid() {
		return errors.New("transaction status is not supported")
	}
	return nil
}

func ValidateName(name, fieldName string) error {
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

func ValidatePhoneNumber(phone string) error {
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

func ValidateSSN(ssn string) error {
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

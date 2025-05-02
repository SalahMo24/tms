package types

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

var allTransactionTypes = []TransactionType{Debit, Credit}

// String returns the string representation
func (t TransactionType) String() string {
	return string(t)
}

// Valid checks if the value is a valid TransactionType
func (t TransactionType) Valid() bool {
	return slices.Contains(allTransactionTypes, t)
}

// Parse parses a string into a TransactionType
func Parse(s string) (TransactionType, error) {
	val := TransactionType(strings.ToLower(s))
	if !val.Valid() {
		return "", fmt.Errorf("invalid TransactionType: %s", s)
	}
	return val, nil
}

// MustParse parses a string and panics if invalid
func MustParse(s string) TransactionType {
	val, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return val
}

// MarshalJSON implements json.Marshaler
func (t TransactionType) MarshalJSON() ([]byte, error) {
	if !t.Valid() {
		return nil, fmt.Errorf("cannot marshal invalid TransactionType: %s", t)
	}
	return json.Marshal(t.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	val, err := Parse(s)
	if err != nil {
		return err
	}
	*t = val
	return nil
}

// Values returns all possible TransactionType values
func Values() []TransactionType {
	return allTransactionTypes
}

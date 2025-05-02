package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Status represents the status of a transaction or process
type Status string

const (
	Pending   Status = "pending"
	Processed Status = "processed"
	Rejected  Status = "rejected"
)

// AllStatuses returns all possible status values
func AllStatuses() []Status {
	return []Status{Pending, Processed, Rejected}
}

// String returns the string representation of the status
func (s Status) String() string {
	return string(s)
}

// Valid checks if the status is valid
func (s Status) Valid() bool {
	switch s {
	case Pending, Processed, Rejected:
		return true
	default:
		return false
	}
}

// ParseStatus converts a string to Status (case-insensitive)
func ParseStatus(str string) (Status, error) {
	val := Status(strings.ToLower(str))
	if !val.Valid() {
		return "", fmt.Errorf("invalid status: %s", str)
	}
	return val, nil
}

// MustParseStatus parses a string to Status or panics if invalid
func MustParseStatus(str string) Status {
	status, err := ParseStatus(str)
	if err != nil {
		panic(err)
	}
	return status
}

// MarshalJSON implements json.Marshaler
func (s Status) MarshalJSON() ([]byte, error) {
	if !s.Valid() {
		return nil, fmt.Errorf("cannot marshal invalid status: %s", s)
	}
	return json.Marshal(s.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Status) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	status, err := ParseStatus(str)
	if err != nil {
		return err
	}
	*s = status
	return nil
}

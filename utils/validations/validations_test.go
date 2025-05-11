package validations

import (
	"testing"
	"tms/app/types"
)

func TestValidateTransactionAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  float64
		wantErr bool
	}{
		{"Positive amount", 100.50, false},
		{"Zero amount", 0, false},
		{"Negative amount", -10.25, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransactionAmount(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransactionAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTransactionType(t *testing.T) {
	tests := []struct {
		name            string
		transactionType types.TransactionType
		wantErr         bool
	}{
		{"Valid type", types.TransactionType("debit"), false},
		{"Invalid type", types.TransactionType("invalid"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransactionType(tt.transactionType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransactionType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTransactionStatus(t *testing.T) {
	tests := []struct {
		name    string
		status  types.Status
		wantErr bool
	}{
		{"Valid status", types.Status("processed"), false},
		{"Valid status", types.Status("pending"), false},
		{"Valid status", types.Status("rejected"), false},

		{"Invalid status", types.Status("invalid"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTransactionStatus(tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransactionStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		field   string
		wantErr bool
	}{
		{"Valid name", "John Doe", "Name", false},
		{"Too short", "Jo", "Name", true},
		{"Too long", "ThisNameIsWayTooLongAndExceedsTheMaximumAllowedLength", "Name", true},
		{"Contains numbers", "John123", "Name", true},
		{"Contains special chars", "John@Doe", "Name", true},
		{"Valid with spaces", "John Michael Doe", "Name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input, tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr bool
	}{
		{"Valid 010", "01012345678", false},
		{"Valid 011", "01112345678", false},
		{"Valid 012", "01212345678", false},
		{"Valid 014", "01412345678", false},
		{"Invalid prefix", "01512345678", true},
		{"Too short", "01012345", true},
		{"Contains letters", "0101234abcd", true},
		{"Contains special chars", "0101234-5678", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhoneNumber(tt.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSSN(t *testing.T) {
	tests := []struct {
		name    string
		ssn     string
		wantErr bool
	}{
		{"Valid SSN", "12345678901234", false},
		{"Too short", "1234567890123", true},
		{"Too long", "123456789012345", true},
		{"Contains letters", "1234567890123a", true},
		{"Contains special chars", "123456789-1234", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSSN(tt.ssn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSSN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

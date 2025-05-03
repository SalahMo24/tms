package accountbalance

import (
	"database/sql"
	"fmt"
	"tms/app/types"
	"tms/internal/database"

	"github.com/shopspring/decimal"
)

type AccountBalanceRepository struct {
	db database.Service
}

func NewRepository() *AccountBalanceRepository {
	return &AccountBalanceRepository{
		db: database.New(), // This gets the singleton database instance
	}
}

func (t *AccountBalanceRepository) Create(tc AccountBalanceCreate) (string, error) {
	db := t.db.DB()

	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Get the last transaction's running balance
	var lastBalance decimal.Decimal
	err = tx.QueryRow(`
        SELECT balance
        FROM account_balance 
        WHERE account_id = $1
		order by created_at desc
		limit 1
		`,
		tc.AccountId,
	).Scan(&lastBalance)

	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("failed to get account balance: %w", err)
	}
	if err == sql.ErrNoRows {
		lastBalance = decimal.Zero
	}

	// Calculate new debit/credit values
	amount := decimal.NewFromFloat(tc.Amount)

	var newBalance decimal.Decimal

	switch tc.TransactionType {
	case types.Debit:

		newBalance = lastBalance.Add(amount)
	case types.Credit:

		newBalance = lastBalance.Sub(amount)
	default:
		return "", fmt.Errorf("invalid transaction type: %s", tc.TransactionType)
	}

	// Verify the new balance won't go negative
	if newBalance.IsNegative() {
		return "", fmt.Errorf("insufficient funds: cannot process %s of %s (current balance: %s)",
			tc.TransactionType, amount.String(), lastBalance.String())
	}

	// Insert new transaction
	var id string
	err = tx.QueryRow(`
        INSERT INTO account_balance (
            account_id, 
            balance,
            transaction_id
        ) VALUES ($1, $2, $3)
        RETURNING id`,
		tc.AccountId,
		newBalance,
		tc.TransactionId,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil
}

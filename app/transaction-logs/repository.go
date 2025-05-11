package transactionlogs

import (
	"tms/app/types"
	"tms/internal/database"
	"tms/utils/assert"
	"tms/utils/validations"
)

type TransactionLogRepository struct {
	db database.Service
}

func NewRepository() *TransactionLogRepository {
	return &TransactionLogRepository{
		db: database.New(), // This gets the singleton database instance
	}
}

func (t *TransactionLogRepository) Create(tl TransactionLogCreate) (string, error) {
	assert.Nil(validations.ValidateTransactionType(tl.TransactionType), "transaction type is not valid")
	assert.Nil(validations.ValidateTransactionAmount(tl.Amount), "transaction amount is not valid")
	assert.Type("", tl.AccountId, "account id should be a string")

	db := t.db.DB()
	query := `
	INSERT INTO transaction_logs (transaction_type, amount, status, account_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`

	var id string
	err := db.QueryRow(
		query,
		tl.TransactionType.String(),
		tl.Amount,
		types.Pending,
		tl.AccountId,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	assert.NotNil(id, "id should not be nil")
	assert.Type("", id, "id should be a string")

	return id, nil
}
func (t *TransactionLogRepository) UpdateStatus(status types.Status, transactionId string) (string, error) {
	assert.Nil(validations.ValidateTransactionStatus(status), "transaction status is not valid")
	assert.Type("", transactionId, "transaction id should be a string")
	db := t.db.DB()

	query := `
	UPDATE transaction_logs SET status=$1 where id = $2
	RETURNING id
`

	var id string
	err := db.QueryRow(
		query,
		status,
		transactionId,
	).Scan(&id)

	if err != nil {
		return "", err
	}
	assert.NotNil(id, "id should not be nil")
	assert.Type("", id, "id should be a string")
	return id, nil
}

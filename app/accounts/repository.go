package accounts

import (
	"tms/internal/database"
	"tms/utils/assert"
)

type AccountRepository struct {
	db database.Service
}

func NewRepository() *AccountRepository {
	return &AccountRepository{
		db: database.New(), // This gets the singleton database instance
	}
}

func (u *AccountRepository) Create(userId string) (string, error) {
	assert.Type("", userId, "user id should be a string")
	db := u.db.DB()

	query := `
	INSERT INTO accounts (user_id, is_active, is_deleted)
	VALUES ($1, $2, $3)
	RETURNING id
`

	var id string
	err := db.QueryRow(
		query,
		userId,
		true,
		false,
	).Scan(&id)

	if err != nil {
		return "", err
	}
	assert.NotNil(id, "account id should not be nil")
	assert.Type("", id, "account id should be a string")

	return id, nil
}

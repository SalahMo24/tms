package accounts

import "tms/internal/database"

type AccountRepository struct {
	db database.Service
}

func NewRepository() *AccountRepository {
	return &AccountRepository{
		db: database.New(), // This gets the singleton database instance
	}
}

func (u *AccountRepository) Create(userId string) (string, error) {
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

	return id, nil
}

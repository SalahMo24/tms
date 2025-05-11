package users

import (
	"tms/internal/database"
	"tms/utils/assert"
	"tms/utils/validations"
)

type UserRepository struct {
	db database.Service
}

func NewRepository() *UserRepository {
	return &UserRepository{
		db: database.New(), // This gets the singleton database instance
	}
}

func (u *UserRepository) Create(user UserCreate) (string, error) {
	assert.Nil(validations.ValidateName(user.FirstName, "first name"), "first name validation failed")
	assert.Nil(validations.ValidateName(user.LastName, "last name"), "last name validation failed")
	assert.Nil(validations.ValidatePhoneNumber(user.PhoneNumber), "Phone number validation failed")
	assert.Nil(validations.ValidateSSN(user.SSN), "SSN validation failed")

	db := u.db.DB()

	query := `
	INSERT INTO users (first_name, last_name, phone_number, ssn)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`

	var id string
	err := db.QueryRow(
		query,

		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.SSN,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	assert.NotNil(id, "user id is nil")
	assert.Type("", id, "id should be an string")

	return id, nil
}
func (u *UserRepository) Exists(ssn string) (string, error) {
	assert.Nil(validations.ValidateSSN(ssn), "SSN validation failed")

	db := u.db.DB()

	query := `
	select id from users where ssn =$1
`

	var id string
	err := db.QueryRow(
		query,
		ssn,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

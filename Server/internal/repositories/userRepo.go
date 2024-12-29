package repositories

import (
	"database/sql"
	"errors"
	"rent-server/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

// constructor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetUserByID(id int) (*models.User, error){
	var user models.User

	err := repo.DB.QueryRow("SELECT id, email, first_name, last_name FROM users where id = $1", id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return nil, errors.New("error getting user")
	}
	
	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.DB.QueryRow("SELECT id, email, first_name, last_name, password_hash FROM users where email = $1", email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user is not registered")
	}
	
	if err != nil {
		return nil, errors.New("error validating user")
	}

	return &user, nil
}

func (repo *UserRepository) CreateUser(data models.User) error {

	_, err := repo.DB.Exec(`INSERT INTO users (email, first_name, last_name, password_hash)
								 VALUES ($1, $2, $3, $4)`, data.Email, data.FirstName, data.LastName, data.PasswordHash)

	if err != nil {
		return errors.New("error registering user")
	}

	return nil
}

func (repo *UserRepository) IsEmailRegistered (email string) (bool, error){
	var exists bool
	err := repo.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
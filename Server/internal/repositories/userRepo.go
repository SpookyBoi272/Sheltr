package repositories

import (
	"database/sql"
	"errors"
	"rent-server/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

//constructor
func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{DB : db}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error){
	var user models.User
	err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user , nil
}

func (r *UserRepository) GetAllUsers() ([]models.User , error){
	rows, err := r.DB.Query("SELECT id, name, email FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next(){
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
package services

import (
	"errors"
	"rent-server/internal/models"
	"rent-server/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

//constructor
func NewUserService(repo *repositories.UserRepository) *UserService{
	return &UserService{Repo: repo}
}

func (s *UserService) GetUser(id int) (*models.User, error){
	return s.Repo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error){
	users, err := s.Repo.GetAllUsers()

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}
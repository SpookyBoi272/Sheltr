package services

import (
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
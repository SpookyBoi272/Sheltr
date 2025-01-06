package services

import (
	"errors"
	"regexp"
	"rent-server/config"
	"rent-server/internal/models"
	"rent-server/internal/repositories"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repositories.UserRepository
	Config *config.Config
}

func NewAuthService(repo *repositories.UserRepository, cfg *config.Config) *AuthService{
	return &AuthService{Repo: repo, Config: cfg}
}

func (s *AuthService) LoginUser(user *models.User) (*models.User, string, error){

	user, err := s.Repo.GetUserByEmail(user.Email)

	if err != nil {
		return nil, "",errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(user.Password))
	if err != nil {
		return nil, "",errors.New("invalid email or password")
	}

	var jwtToken string
	jwtToken, err = generateJWT(user, s.Config.JwtKey)
	if err != nil {
		return nil, "", errors.New("error validating user")
	}

	user.PasswordHash = ""
	return user, jwtToken, nil
}

func (s *AuthService) SignupUser(data *models.User) (*models.User, string, error){
	var err error

	err = validateEmail(data.Email, s)
    if err != nil {
        return nil, "", err
    }


	err = validatePassword(data.Password)
    if err != nil {
        return nil, "", err
    }

    err = validateName(data.FirstName)
    if err != nil {
        return nil, "", err
    }

    err = validateName(data.LastName)
    if err != nil {
        return nil, "", err
    }



	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("invalid password format")
	}
	data.PasswordHash = string(hashedPwd)
	
	err = s.Repo.CreateUser(*data)

	if err != nil {
		return nil, "",err
	}

	var jwtToken string
	jwtToken, err = generateJWT(data, s.Config.JwtKey)
	if err != nil {
		return nil, "", errors.New("error validating user")
	}

	data.PasswordHash = ""
	return data, jwtToken, nil
}

func validateEmail(email string, s *AuthService) error {
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    if !re.MatchString(email) {
        return errors.New("invalid email format")
    }
	
	exists, err := s.Repo.IsEmailRegistered(email)
	if err != nil {
		return errors.New("error registering user")
	} else if exists{
		return errors.New("user already registered")
	}

	return nil
}

func validatePassword(password string) error{
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func validateName(name string) error {
    if name == "" {
        return errors.New("name cannot be blank")
    }
    return nil
}

func generateJWT(user *models.User, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(user.ID),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
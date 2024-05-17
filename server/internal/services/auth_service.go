package services

import (
	"errors"
	"log"
	"strength-forge-app/internal/models"
	"strength-forge-app/internal/repositories"
	"strength-forge-app/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user *models.User) (string, error) {
	_, err := s.repo.GetUserByEmail(user.Email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	user.Password = string(hashedPassword)

	err = s.repo.CreateUser(user)
	if err != nil {
		log.Println(err)
		return "", err
	}

	token, err := utils.CreateToken(user.Email)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (s *AuthService) LogIn(user *models.User) (string, error) {
	userFromDB, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.CreateToken(user.Email)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (s *AuthService) Authenticate(tokenString string) error {
	return utils.VerifyToken(tokenString)
}

package services

import (
	"errors"
	"log"
	"strength-forge-app/internal/dtos"
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

func (s *AuthService) CreateUser(registerUser *dtos.RegisterUser) (string, error) {
	_, err := s.repo.GetUserByEmail(registerUser.Email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var user = &models.User{
		Email:    registerUser.Email,
		Password: string(hashedPassword),
	}

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

func (s *AuthService) LogIn(loginUser *dtos.LoginUser) (string, error) {
	user, err := s.repo.GetUserByEmail(loginUser.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.CreateToken(loginUser.Email)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (s *AuthService) Authenticate(tokenString string) error {
	return utils.VerifyToken(tokenString)
}

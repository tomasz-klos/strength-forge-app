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

type AuthService interface {
	CreateUser(registerUser *dtos.RegisterUser) (string, error)
	LogIn(loginUser *dtos.LoginUser) (string, error)
	Authenticate(tokenString string) error
}

type authService struct {
	repo           repositories.UserRepository
	tokenGenerator utils.TokenGenerator
}

func NewAuthService(repo repositories.UserRepository, tokenGenerator utils.TokenGenerator) AuthService {
	return &authService{
		repo:           repo,
		tokenGenerator: tokenGenerator,
	}
}

func (s *authService) CreateUser(registerUser *dtos.RegisterUser) (string, error) {
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

	token, err := s.tokenGenerator.CreateToken(user.Email)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (s *authService) LogIn(loginUser *dtos.LoginUser) (string, error) {
	user, err := s.repo.GetUserByEmail(loginUser.Email)
	log.Println(user, loginUser.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.tokenGenerator.CreateToken(loginUser.Email)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}

func (s *authService) Authenticate(tokenString string) error {
	return s.tokenGenerator.VerifyToken(tokenString)
}

package services

import (
	"strength-forge-app/internal/dtos"
	"strength-forge-app/internal/repositories"
	"strength-forge-app/utils"
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

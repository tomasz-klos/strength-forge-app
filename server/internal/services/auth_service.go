package services

import (
	"errors"
	"strength-forge-app/internal/models"
	"strength-forge-app/internal/repositories"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *AuthService) LogIn(user *models.User) error {
	userFromDB, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	if userFromDB.Password != user.Password {
		return ErrInvalidCredentials
	}

	return nil
}

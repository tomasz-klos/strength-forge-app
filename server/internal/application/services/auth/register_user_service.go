package services_auth

import (
	"errors"
	"log"
	"strength-forge-app/internal/domain/dtos"
	"strength-forge-app/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Register(registerUser *dtos.RegisterUser) (string, error) {
	_, err := s.repo.GetUserByEmail(registerUser.Email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
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

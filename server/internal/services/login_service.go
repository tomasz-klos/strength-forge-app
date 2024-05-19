package services

import (
	"errors"
	"log"
	"strength-forge-app/internal/dtos"

	"golang.org/x/crypto/bcrypt"
)

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

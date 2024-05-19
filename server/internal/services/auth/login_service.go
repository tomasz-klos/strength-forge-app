package services_auth

import (
	"errors"
	"strength-forge-app/internal/dtos"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) LogIn(loginUser *dtos.LoginUser) (string, error) {
	user, err := s.repo.GetUserByEmail(loginUser.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.tokenGenerator.CreateToken(loginUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

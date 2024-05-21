package services_auth

import (
	"errors"
	"log"
	"strength-forge-app/internal/domain/dtos"
)

func (s *authService) ValidateToken(tokenString string) (*dtos.ResponseUser, error) {
	email, err := s.tokenGenerator.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	responseUser := &dtos.ResponseUser{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	return responseUser, nil
}

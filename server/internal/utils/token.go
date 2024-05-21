package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator interface {
	CreateToken(email string) (string, error)
	VerifyToken(tokenString string) (string, error)
}

type tokenGenerator struct{}

func NewTokenGenerator() TokenGenerator {
	return &tokenGenerator{}
}

func (g *tokenGenerator) CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey, err := GetEnvVariable("SECRET_KEY")
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (g *tokenGenerator) VerifyToken(tokenString string) (string, error) {
	secretKey, err := GetEnvVariable("SECRET_KEY")
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims)["email"].(string), nil
}

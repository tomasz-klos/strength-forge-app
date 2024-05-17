package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email string) (string, error) {
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

func VerifyToken(tokenString string) error {
	secretKey, err := GetEnvVariable("SECRET_KEY")
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

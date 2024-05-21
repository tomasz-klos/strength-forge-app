package repositories

import (
	"strength-forge-app/internal/domain/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

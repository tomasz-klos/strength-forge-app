package db

import (
	"strength-forge-app/internal/domain/models"
)

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
}

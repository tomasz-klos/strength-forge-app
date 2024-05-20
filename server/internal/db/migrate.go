package db

import (
	"strength-forge-app/internal/models"
)

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
}

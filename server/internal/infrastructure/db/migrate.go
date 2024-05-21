package db

import (
	"strength-forge-app/internal/domain/models"
)

func AutoMigrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}

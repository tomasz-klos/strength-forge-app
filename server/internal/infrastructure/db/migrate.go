package db

import (
	"log"
	"strength-forge-app/internal/domain/models"
)

func AutoMigrate() {
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}
}

package db

import (
	"fmt"
	"log"
	"strength-forge-app/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDatabaseConfig() (map[string]string, error) {
	keys := []string{"DB_NAME", "DB_USER", "DB_PASSWORD", "DB_HOST"}
	config := make(map[string]string)

	for _, key := range keys {
		value, err := utils.GetEnvVariable(key)
		if err != nil {
			return nil, fmt.Errorf("failed to get env variable %s: %v", key, err)
		}
		config[key] = value
	}

	return config, nil
}

func Init() {
	var err error

	dbConfig, err := getDatabaseConfig()
	if err != nil {
		log.Fatalf("error getting database config: %v", err)
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConfig["DB_USER"], dbConfig["DB_PASSWORD"], dbConfig["DB_HOST"], dbConfig["DB_NAME"])

	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
}

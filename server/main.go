package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"strength-forge-app/utils"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	dbName, err := utils.GetEnvVariable("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbUser, err := utils.GetEnvVariable("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := utils.GetEnvVariable("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbHost, err := utils.GetEnvVariable("DB_HOST")
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	var err error
	mux := http.NewServeMux()
	corsHandler := cors.Default().Handler(mux)

	db, err = ConnectDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

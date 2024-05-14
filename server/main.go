package main

import (
	"log"
	"net/http"

	"strength-forge-app/db"

	"github.com/rs/cors"
)

func main() {
	var err error
	mux := http.NewServeMux()
	corsHandler := cors.Default().Handler(mux)

	db, err := db.Init()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	log.Println("Server is running on port 8080")
	log.Print("Connected to database: ", db.Migrator().CurrentDatabase())
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

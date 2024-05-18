package main

import (
	"log"
	"net/http"

	"strength-forge-app/db"
	"strength-forge-app/handlers"
	"strength-forge-app/internal/repositories"
	"strength-forge-app/internal/services"

	"github.com/rs/cors"
)

func init() {
	db.Init()
	db.AutoMigrate()
}

func main() {
	mux := http.NewServeMux()
	corsHandler := cors.Default().Handler(mux)
	userRepo := repositories.NewPostgresUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	userHandler := handlers.NewAuthHandler(*authService)

	mux.HandleFunc("/api/auth/validate-token", userHandler.ValidateToken)
	mux.HandleFunc("/api/auth/register", userHandler.RegisterUser)
	mux.HandleFunc("/api/auth/login", userHandler.SignIn)
	mux.HandleFunc("/api/auth/logout", userHandler.SignOut)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

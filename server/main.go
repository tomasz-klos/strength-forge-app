package main

import (
	"log"
	"net/http"

	"strength-forge-app/db"
	"strength-forge-app/handlers"
	"strength-forge-app/internal/repositories"
	"strength-forge-app/internal/services"
	"strength-forge-app/utils"

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
	authService := services.NewAuthService(userRepo, utils.NewTokenGenerator())
	authHandler := handlers.NewAuthHandler(authService)

	mux.HandleFunc("/api/auth/validate-token", authHandler.ValidateToken)
	mux.HandleFunc("/api/auth/register", authHandler.RegisterUser)
	mux.HandleFunc("/api/auth/signin", authHandler.SignIn)
	mux.HandleFunc("/api/auth/signout", authHandler.SignOut)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

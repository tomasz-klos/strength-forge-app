package main

import (
	"log"
	"net/http"

	auth_handler "strength-forge-app/handlers/auth"
	auth_service "strength-forge-app/internal/application/services/auth"
	"strength-forge-app/internal/infrastructure/db"
	"strength-forge-app/internal/infrastructure/repositories"
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
	authService := auth_service.NewAuthService(userRepo, utils.NewTokenGenerator())
	authHandler := auth_handler.NewAuthHandler(authService)

	mux.HandleFunc("/api/auth/validate-token", authHandler.ValidateToken)
	mux.HandleFunc("/api/auth/register", authHandler.RegisterUser)
	mux.HandleFunc("/api/auth/signin", authHandler.SignIn)
	mux.HandleFunc("/api/auth/signout", authHandler.SignOut)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

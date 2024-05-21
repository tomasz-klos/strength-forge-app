package handlers_auth

import (
	services_auth "strength-forge-app/internal/application/services/auth"
)

type AuthHandler struct {
	service services_auth.AuthService
}

func NewAuthHandler(service services_auth.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

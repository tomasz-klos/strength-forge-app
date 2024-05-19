package handlers_auth

import (
	auth_service "strength-forge-app/internal/services/auth"
)

type AuthHandler struct {
	service auth_service.AuthService
}

func NewAuthHandler(service auth_service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

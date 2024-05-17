package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strength-forge-app/internal/dtos"
	"strength-forge-app/internal/services"
	"time"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUser dtos.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&registerUser)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	token, err := h.service.CreateUser(&registerUser)
	if err != nil {
		if err.Error() == "user already exists" {
			writeJSONResponse(w, http.StatusConflict, JSONResponse{Error: MsgUserAlreadyExists})
			return
		}

		log.Println(err)
		writeJSONResponse(w, http.StatusInternalServerError, JSONResponse{Error: MsgInternalError})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	writeJSONResponse(w, http.StatusCreated, JSONResponse{Message: MsgUserCreated})
}

func (h *AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var loginUser dtos.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	token, err := h.service.LogIn(&loginUser)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgUnauthorized})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: MsgLoginSuccessful})
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: "Logged out"})
}

func (h *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Cookies()[0].Value
	log.Println(tokenString)
	if tokenString == "" {
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgNoTokenProvided})
		return
	}

	err := h.service.Authenticate(tokenString)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgUnauthorized})
		return
	}

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: "Authenticated"})
}

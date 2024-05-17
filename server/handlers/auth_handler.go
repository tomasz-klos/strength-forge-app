package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strength-forge-app/internal/models"
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
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	token, err := h.service.CreateUser(&user)
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
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	token, err := h.service.LogIn(&user)
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

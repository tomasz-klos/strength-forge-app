package handlers_auth

import (
	"encoding/json"
	"log"
	"net/http"
	"strength-forge-app/internal/dtos"
	"time"
)

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, JSONResponse{Error: MsgMethodNotAllowed})
		return
	}

	var registerUser dtos.RegisterUser

	err := json.NewDecoder(r.Body).Decode(&registerUser)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	token, err := h.service.Register(&registerUser)
	if err.Error() == "user already exists" {
		writeJSONResponse(w, http.StatusConflict, JSONResponse{Error: MsgUserAlreadyExists})
		return
	}

	if err != nil {
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
		Path:     "/",
	})

	writeJSONResponse(w, http.StatusCreated, JSONResponse{Message: MsgUserCreated})
}

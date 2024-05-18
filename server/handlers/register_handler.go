package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strength-forge-app/internal/dtos"
	"time"
)

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

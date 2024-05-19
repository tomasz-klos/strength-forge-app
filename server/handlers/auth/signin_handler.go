package handlers_auth

import (
	"encoding/json"
	"log"
	"net/http"
	"strength-forge-app/internal/dtos"
	"time"
)

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
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
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgInvalidPayload})
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

package handlers_auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strength-forge-app/internal/dtos"
	"time"
)

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, JSONResponse{Error: MsgMethodNotAllowed})
		return
	}

	var loginUser dtos.LoginUser

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, JSONResponse{Error: MsgInvalidPayload})
		return
	}

	token, err := h.service.SignIn(&loginUser)
	if err != nil {
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgUnauthorized})
		return
	}

	if token == "" {
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

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: MsgLoginSuccessful})
}

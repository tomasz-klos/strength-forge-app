package handlers_auth

import (
	"net/http"
)

func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, JSONResponse{Error: MsgMethodNotAllowed})
		return
	}

	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" {
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgNoTokenProvided})
		return
	}

	tokenString := cookie.Value

	user, err := h.service.ValidateToken(tokenString)
	if err != nil {

		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgUnauthorized})
		return
	}

	if user == nil {
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgUnauthorized})
		return
	}

	writeJSONResponse(w, http.StatusOK, JSONResponse{Data: user})
}

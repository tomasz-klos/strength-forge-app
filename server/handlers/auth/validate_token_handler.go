package handlers_auth

import (
	"log"
	"net/http"
)

func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Cookies()[0].Value
	log.Println(tokenString)
	if tokenString == "" {
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgNoTokenProvided})
		return
	}

	err := h.service.ValidateToken(tokenString)
	if err != nil {
		log.Println(err)
		writeJSONResponse(w, http.StatusUnauthorized, JSONResponse{Error: MsgUnauthorized})
		return
	}

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: "Authenticated"})
}

package handlers_auth

import (
	"net/http"
	"time"
)

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: "Logged out"})
}

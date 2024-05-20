package handlers_auth

import (
	"net/http"
	"time"
)

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONResponse(w, http.StatusMethodNotAllowed, JSONResponse{Error: MsgMethodNotAllowed})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	writeJSONResponse(w, http.StatusOK, JSONResponse{Message: MsgLogoutSuccessful})
}

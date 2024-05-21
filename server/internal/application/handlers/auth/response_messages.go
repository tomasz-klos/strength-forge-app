package handlers_auth

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

const (
	MsgLoginSuccessful   = "login successful"
	MsgLogoutSuccessful  = "logout successful"
	MsgUserCreated       = "user created successfully"
	MsgInvalidPayload    = "invalid request payload"
	MsgUserAlreadyExists = "user already exists"
	MsgInternalError     = "internal server error"
	MsgUnauthorized      = "unauthorized"
	MsgMethodNotAllowed  = "method not allowed"
	MsgNoTokenProvided   = "no token provided"
)

func writeJSONResponse(w http.ResponseWriter, status int, response JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

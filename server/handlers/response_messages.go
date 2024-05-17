package handlers

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
	MsgLoginSuccessful   = "Login successful"
	MsgUserCreated       = "User created successfully"
	MsgInvalidPayload    = "Invalid request payload"
	MsgUserAlreadyExists = "User already exists"
	MsgInternalError     = "Internal server error"
	MsgUnauthorized      = "Unauthorized"
	MsgMethodNotAllowed  = "Method not allowed"
	MsgNoTokenProvided   = "No token provided"
)

func writeJSONResponse(w http.ResponseWriter, status int, response JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

package handlers_auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	helpers_test "strength-forge-app/helpers"
	handlers_auth "strength-forge-app/internal/application/handlers/auth"
	mock_services "strength-forge-app/internal/application/services/auth/mock"
	"strength-forge-app/internal/domain/dtos"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockAuthService(ctrl)

	loginUser := dtos.LoginUser{
		Email:    "test1@example.com",
		Password: "password123",
	}

	handler := http.HandlerFunc(handlers_auth.NewAuthHandler(mockService).SignIn)

	testCases := []struct {
		name             string
		method           string
		input            interface{}
		mockSetup        func()
		expectedStatus   int
		expectedResponse handlers_auth.JSONResponse
	}{
		{
			name:   "Success",
			method: "POST",
			input:  loginUser,
			mockSetup: func() {
				mockService.EXPECT().SignIn(&loginUser).Return("token", nil)
			},
			expectedStatus: http.StatusOK,
			expectedResponse: handlers_auth.JSONResponse{
				Message: handlers_auth.MsgLoginSuccessful,
			},
		},
		{
			name:           "Invalid Payload",
			method:         "POST",
			input:          "invalid payload",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgInvalidPayload,
			},
		},
		{
			name:           "Read body error",
			method:         "POST",
			input:          "",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgInvalidPayload,
			},
		},
		{
			name:   "Unauthorized",
			method: "POST",
			input:  loginUser,
			mockSetup: func() {
				mockService.EXPECT().SignIn(&loginUser).Return("", errors.New("unauthorized"))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgUnauthorized,
			},
		},
		{
			name:   "Token is empty",
			method: "POST",
			input:  loginUser,
			mockSetup: func() {
				mockService.EXPECT().SignIn(&loginUser).Return("", nil)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgInternalError,
			},
		},
		{
			name:   "No Token Provided",
			method: "POST",
			input:  loginUser,
			mockSetup: func() {
				mockService.EXPECT().SignIn(&loginUser).Return("", errors.New("no token provided"))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgUnauthorized,
			},
		},
		{
			name:           "Method not allowed",
			method:         "GET",
			input:          nil,
			mockSetup:      func() {},
			expectedStatus: http.StatusMethodNotAllowed,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgMethodNotAllowed,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body io.Reader

			switch input := tc.input.(type) {
			case string:
				if input == "" && tc.name == "Read body error" {
					body = &helpers_test.ErrorReader{}
				} else {
					body = bytes.NewBuffer([]byte(input))
				}
			case dtos.LoginUser:
				payload, _ := json.Marshal(input)
				body = bytes.NewBuffer(payload)
			}

			tc.mockSetup()

			recorder := helpers_test.ExecuteRequest(t, tc.method, "/api/auth/signin", body, handler)

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			var response handlers_auth.JSONResponse
			err := json.NewDecoder(recorder.Body).Decode(&response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}

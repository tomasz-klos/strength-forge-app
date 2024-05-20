package handlers_auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	handlers_auth "strength-forge-app/handlers/auth"
	helpers_test "strength-forge-app/helpers"
	"strength-forge-app/internal/domain/dtos"
	mock_services "strength-forge-app/internal/services/auth/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockAuthService(ctrl)

	registerUser := dtos.RegisterUser{
		Name:     "TestUser",
		Email:    "test1@example.com",
		Password: "password123",
	}

	handler := http.HandlerFunc(handlers_auth.NewAuthHandler(mockService).RegisterUser)

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
			input:  registerUser,
			mockSetup: func() {
				mockService.EXPECT().Register(&registerUser).Return("token", nil)
			},
			expectedStatus: http.StatusCreated,
			expectedResponse: handlers_auth.JSONResponse{
				Message: handlers_auth.MsgUserCreated,
			},
		},
		{
			name:   "User already exists",
			method: "POST",
			input:  registerUser,
			mockSetup: func() {
				mockService.EXPECT().Register(&registerUser).Return("", errors.New("user already exists"))
			},
			expectedStatus: http.StatusConflict,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgUserAlreadyExists,
			},
		},
		{
			name:   "Error",
			method: "POST",
			input:  registerUser,
			mockSetup: func() {
				mockService.EXPECT().Register(&registerUser).Return("", errors.New("error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgInternalError,
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
			name:           "Invalid payload",
			method:         "POST",
			input:          "invalid",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgInvalidPayload,
			},
		},
		{
			name:   "Token is empty",
			method: "POST",
			input:  registerUser,
			mockSetup: func() {
				mockService.EXPECT().Register(&registerUser).Return("", nil)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgInternalError,
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
			case dtos.RegisterUser:
				payload, _ := json.Marshal(input)
				body = bytes.NewBuffer(payload)
			}

			tc.mockSetup()

			recorder := helpers_test.ExecuteRequest(t, tc.method, "/api/auth/register", body, handler)

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			if tc.expectedStatus == http.StatusOK || tc.expectedStatus == http.StatusCreated {
				var jsonResponse handlers_auth.JSONResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &jsonResponse)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, jsonResponse)

				if tc.name == "Success" {
					cookies := recorder.Result().Cookies()
					var tokenCookie *http.Cookie
					for _, cookie := range cookies {
						if cookie.Name == "token" {
							tokenCookie = cookie
							break
						}
					}
					assert.NotNil(t, tokenCookie)
					assert.Equal(t, "token", tokenCookie.Name)
					assert.Equal(t, "token", tokenCookie.Value)
					assert.Equal(t, http.SameSiteStrictMode, tokenCookie.SameSite)
					assert.True(t, tokenCookie.HttpOnly)
				}
			} else {
				var jsonResponse handlers_auth.JSONResponse
				err := json.Unmarshal(recorder.Body.Bytes(), &jsonResponse)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, jsonResponse)
			}
		})
	}
}

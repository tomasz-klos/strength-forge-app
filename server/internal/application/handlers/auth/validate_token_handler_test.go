package handlers_auth_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers_auth "strength-forge-app/internal/application/handlers/auth"
	mock_services "strength-forge-app/internal/application/services/auth/mock"
	"strength-forge-app/internal/domain/dtos"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_ValidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_services.NewMockAuthService(ctrl)

	testToken := "valid_token"
	testUser := &dtos.ResponseUser{
		Id:    uuid.New(),
		Name:  "TestUser",
		Email: "test1@example.com",
	}

	handler := http.HandlerFunc(handlers_auth.NewAuthHandler(mockService).ValidateToken)

	testCases := []struct {
		name             string
		method           string
		token            string
		mockSetup        func()
		expectedStatus   int
		expectedResponse handlers_auth.JSONResponse
	}{
		{
			name:   "Success",
			method: "GET",
			token:  testToken,
			mockSetup: func() {
				mockService.EXPECT().ValidateToken(testToken).Return(testUser, nil).Times(1)
			},
			expectedStatus: http.StatusOK,
			expectedResponse: handlers_auth.JSONResponse{
				Data: map[string]interface{}{
					"id":    testUser.Id.String(),
					"name":  testUser.Name,
					"email": testUser.Email,
				},
			},
		},
		{
			name:   "No token provided or empty token",
			method: "GET",
			token:  "",
			mockSetup: func() {
				mockService.EXPECT().ValidateToken(gomock.Any()).Times(0)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgNoTokenProvided,
			},
		},
		{
			name:   "Error",
			method: "GET",
			token:  testToken,
			mockSetup: func() {
				mockService.EXPECT().ValidateToken(testToken).Return(nil, errors.New("error")).Times(1)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgUnauthorized,
			},
		},
		{
			name:   "User not found",
			method: "GET",
			token:  testToken,
			mockSetup: func() {
				mockService.EXPECT().ValidateToken(testToken).Return(nil, nil).Times(1)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgUnauthorized,
			},
		},
		{
			name:   "Method not allowed",
			method: "POST",
			token:  testToken,
			mockSetup: func() {
				mockService.EXPECT().ValidateToken(testToken).Times(0)
			},
			expectedStatus: http.StatusMethodNotAllowed,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgMethodNotAllowed,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body io.Reader

			req, err := http.NewRequest("GET", "api/auth/validate-token", body)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")

			req.Method = tc.method

			if tc.token != "" || tc.name == "No token provided or empty token" {
				cookie := &http.Cookie{
					Name:  "token",
					Value: tc.token,
				}
				req.AddCookie(cookie)
			}

			tc.mockSetup()

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			var jsonResponse handlers_auth.JSONResponse
			err = json.Unmarshal(recorder.Body.Bytes(), &jsonResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, jsonResponse)
		})
	}
}

package handlers_auth_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	helpers_test "strength-forge-app/helpers"
	handlers_auth "strength-forge-app/internal/application/handlers/auth"

	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_SignOut(t *testing.T) {
	testToken := "valid_token"
	handler := http.HandlerFunc(handlers_auth.NewAuthHandler(nil).SignOut)

	testCases := []struct {
		name             string
		method           string
		token            string
		expectedStatus   int
		expectedResponse handlers_auth.JSONResponse
	}{
		{
			name:           "Success",
			method:         "POST",
			token:          testToken,
			expectedStatus: http.StatusOK,
			expectedResponse: handlers_auth.JSONResponse{
				Message: handlers_auth.MsgLogoutSuccessful,
			},
		},
		{
			name:           "Method Not Allowed",
			method:         "GET",
			token:          testToken,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedResponse: handlers_auth.JSONResponse{
				Error: handlers_auth.MsgMethodNotAllowed,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := helpers_test.ExecuteRequest(t, tc.method, "/api/auth/signout", nil, handler)

			assert.Equal(t, tc.expectedStatus, recorder.Code)

			var response handlers_auth.JSONResponse
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expectedResponse, response)

			if recorder.Code != http.StatusMethodNotAllowed {
				cookies := recorder.Result().Cookies()
				var tokenCookie *http.Cookie
				for _, cookie := range cookies {
					if cookie.Name == "token" {
						tokenCookie = cookie
						break
					}
				}
				assert.NotNil(t, tokenCookie)
				assert.Equal(t, "", tokenCookie.Value)
				assert.True(t, tokenCookie.Expires.Before(time.Now()))
			}
		})
	}
}

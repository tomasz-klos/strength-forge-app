package services_auth_test

import (
	"errors"
	"strength-forge-app/internal/dtos"
	"strength-forge-app/internal/models"
	mock_repositories "strength-forge-app/internal/repositories/mock"
	services "strength-forge-app/internal/services/auth"
	mock_utils "strength-forge-app/utils/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_repositories.NewMockUserRepository(ctrl)
	tokenGeneratorMock := mock_utils.NewMockTokenGenerator(ctrl)
	authService := services.NewAuthService(userRepoMock, tokenGeneratorMock)

	loginUser := &dtos.LoginUser{
		Email:    "test@email.com",
		Password: "password123",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(loginUser.Password), bcrypt.DefaultCost)
	expectedUser := &models.User{
		Email:    loginUser.Email,
		Password: string(hashedPassword),
	}

	expectedUserWrongPassword := &models.User{
		Email:    loginUser.Email,
		Password: "wrong_password",
	}

	testCases := []struct {
		name     string
		mockFunc func()
		expected string
		err      error
	}{
		{
			name: "User not found",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(nil, errors.New("not found")).Times(1)
			},
			expected: "",
			err:      errors.New("not found"),
		},
		{
			name: "Invalid credentials",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(expectedUserWrongPassword, nil).Times(1)

			},
			expected: "",
			err:      errors.New("invalid credentials"),
		},
		{
			name: "Error creating token",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(expectedUser, nil).Times(1)
				tokenGeneratorMock.EXPECT().CreateToken(loginUser.Email).Return("", errors.New("create error")).Times(1)
			},
			expected: "",
			err:      errors.New("create error"),
		},
		{
			name: "Success",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(expectedUser, nil).Times(1)
				tokenGeneratorMock.EXPECT().CreateToken(loginUser.Email).Return("expected_token", nil).Times(1)
			},
			expected: "expected_token",
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			token, err := authService.LogIn(loginUser)
			assert.Equal(t, tc.expected, token)
			assert.Equal(t, tc.err, err)
		})
	}
}

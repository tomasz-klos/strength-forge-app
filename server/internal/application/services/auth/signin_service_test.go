package services_auth_test

import (
	"errors"
	services "strength-forge-app/internal/application/services/auth"
	"strength-forge-app/internal/domain/dtos"
	"strength-forge-app/internal/domain/models"
	mock_repositories "strength-forge-app/internal/infrastructure/repositories/mock"
	mock_utils "strength-forge-app/utils/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_SignIn(t *testing.T) {
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

	testCases := []struct {
		name     string
		mockFunc func()
		expected string
		err      error
	}{
		{
			name: "Error getting user by email",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(nil, errors.New("error getting user by email")).Times(1)
			},
			expected: "",
			err:      errors.New("error getting user by email"),
		},
		{
			name: "User is nil",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(nil, nil).Times(1)
			},
			expected: "",
			err:      errors.New("user not found"),
		},
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
				expectedUser := &models.User{
					Email:    loginUser.Email,
					Password: "invalid_password",
				}
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(expectedUser, nil).Times(1)

			},
			expected: "",
			err:      errors.New("invalid credentials"),
		},
		{
			name: "Error creating token",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(loginUser.Email).Return(expectedUser, nil).Times(1)
				tokenGeneratorMock.EXPECT().CreateToken(loginUser.Email).Return("", errors.New("error creating token")).Times(1)
			},
			expected: "",
			err:      errors.New("error creating token"),
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
			token, err := authService.SignIn(loginUser)
			assert.Equal(t, tc.expected, token)
			assert.Equal(t, tc.err, err)
		})
	}
}

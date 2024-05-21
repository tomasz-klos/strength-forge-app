package services_auth_test

import (
	"errors"
	"testing"

	"strength-forge-app/internal/domain/dtos"
	"strength-forge-app/internal/domain/models"
	mock_repositories "strength-forge-app/internal/infrastructure/repositories/mock"
	services "strength-forge-app/internal/services/auth"
	mock_utils "strength-forge-app/utils/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_repositories.NewMockUserRepository(ctrl)
	tokenGeneratorMock := mock_utils.NewMockTokenGenerator(ctrl)
	authService := services.NewAuthService(userRepoMock, tokenGeneratorMock)

	registerUser := &dtos.RegisterUser{
		Email:    "test@example.com",
		Password: "password123",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	expectedUser := &models.User{
		Email:    registerUser.Email,
		Password: string(hashedPassword),
	}

	testCases := []struct {
		name     string
		mockFunc func()
		expected string
		err      error
	}{
		{
			name: "User already exists",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(registerUser.Email).Return(expectedUser, nil).Times(1)
			},
			expected: "",
			err:      errors.New("user already exists"),
		},
		{
			name: "Error creating user",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(registerUser.Email).Return(nil, errors.New("not found")).Times(1)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(errors.New("create error")).Times(1)
			},
			expected: "",
			err:      errors.New("create error"),
		},
		{
			name: "Error creating token",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(registerUser.Email).Return(nil, errors.New("not found")).Times(1)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil).Times(1)
				tokenGeneratorMock.EXPECT().CreateToken(registerUser.Email).Return("", errors.New("token error")).Times(1)
			},
			expected: "",
			err:      errors.New("token error"),
		},
		{
			name: "Success",
			mockFunc: func() {
				userRepoMock.EXPECT().GetUserByEmail(registerUser.Email).Return(nil, errors.New("not found")).Times(1)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil).Times(1)
				tokenGeneratorMock.EXPECT().CreateToken(registerUser.Email).Return("expected_token", nil).Times(1)
			},
			expected: "expected_token",
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			token, err := authService.Register(registerUser)
			assert.Equal(t, tc.expected, token)
			assert.Equal(t, tc.err, err)
		})
	}
}

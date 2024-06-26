package services_auth_test

import (
	"errors"
	services_auth "strength-forge-app/internal/application/services/auth"
	"strength-forge-app/internal/domain/dtos"
	"strength-forge-app/internal/domain/models"
	mock_repositories "strength-forge-app/internal/infrastructure/repositories/mock"
	mock_utils "strength-forge-app/internal/utils/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_ValidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_repositories.NewMockUserRepository(ctrl)
	tokenGeneratorMock := mock_utils.NewMockTokenGenerator(ctrl)
	authService := services_auth.NewAuthService(userRepoMock, tokenGeneratorMock)

	token := "expected_token"
	email := "test@email.com"
	testUser := &models.User{
		Id:    uuid.Nil,
		Name:  "Test User",
		Email: email,
	}

	testCases := []struct {
		name         string
		mockFunc     func()
		expectedUser *dtos.ResponseUser
		expectedErr  error
	}{
		{
			name: "Invalid token",
			mockFunc: func() {
				tokenGeneratorMock.EXPECT().VerifyToken(token).Return("", errors.New("invalid token")).Times(1)
			},
			expectedUser: nil,
			expectedErr:  errors.New("invalid token"),
		},
		{
			name: "User not found",
			mockFunc: func() {
				tokenGeneratorMock.EXPECT().VerifyToken(token).Return(email, nil).Times(1)
				userRepoMock.EXPECT().GetUserByEmail(email).Return(nil, errors.New("user not found")).Times(1)
			},
			expectedUser: nil,
			expectedErr:  errors.New("user not found"),
		},
		{
			name: "User is nil",
			mockFunc: func() {
				tokenGeneratorMock.EXPECT().VerifyToken(token).Return(email, nil).Times(1)
				userRepoMock.EXPECT().GetUserByEmail(email).Return(nil, nil).Times(1)
			},
			expectedUser: nil,
			expectedErr:  errors.New("user not found"),
		},
		{
			name: "User found",
			mockFunc: func() {
				tokenGeneratorMock.EXPECT().VerifyToken(token).Return(email, nil).Times(1)
				userRepoMock.EXPECT().GetUserByEmail(email).Return(testUser, nil).Times(1)
			},
			expectedUser: &dtos.ResponseUser{
				Id:    testUser.Id,
				Name:  testUser.Name,
				Email: testUser.Email,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			user, err := authService.ValidateToken(token)
			assert.Equal(t, tc.expectedUser, user)
			assert.Equal(t, tc.expectedErr, err)
		})
	}

}

package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestAuthenticateUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should authenticate", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		token := "token_test"
		userRepo := mock.NewMockUserRepositoryInterface(ctrl)
		userRepo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

		crypt := mock.NewMockCryptographyInterface(ctrl)
		crypt.EXPECT().Encrypt(gomock.Any()).Return(&token, nil).AnyTimes()

		useCase := uc.NewDefaulAuthenticateUseCase(userRepo, hasher, crypt)

		input := uc.AuthenticateUseCaseInput{
			Email:    "test@mail.com",
			Password: "Test@123",
		}

		result, err := useCase.Execute(input)

		require.Equal(t, false, len(err) > 0)
		require.NotNil(t, result)
	})

	t.Run("should not authenticate when email is invalid", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		token := "token_test"
		userRepo := mock.NewMockUserRepositoryInterface(ctrl)
		userRepo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

		crypt := mock.NewMockCryptographyInterface(ctrl)
		crypt.EXPECT().Encrypt(gomock.Any()).Return(&token, nil).AnyTimes()

		useCase := uc.NewDefaulAuthenticateUseCase(userRepo, hasher, crypt)

		input := uc.AuthenticateUseCaseInput{
			Email:    "invalid_email",
			Password: "Test@123",
		}

		result, err := useCase.Execute(input)

		require.Equal(t, true, len(err) > 0)
		require.Nil(t, result)
	})

	t.Run("should not authenticate when email is invalid", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		token := "token_test"
		userRepo := mock.NewMockUserRepositoryInterface(ctrl)
		userRepo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

		crypt := mock.NewMockCryptographyInterface(ctrl)
		crypt.EXPECT().Encrypt(gomock.Any()).Return(&token, nil).AnyTimes()

		useCase := uc.NewDefaulAuthenticateUseCase(userRepo, hasher, crypt)

		input := uc.AuthenticateUseCaseInput{
			Email:    "invalid_email",
			Password: "Test@123",
		}

		result, err := useCase.Execute(input)

		require.Equal(t, true, len(err) > 0)
		require.Nil(t, result)
	})

	t.Run("should not authenticate when password is invalid", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		token := "token_test"
		userRepo := mock.NewMockUserRepositoryInterface(ctrl)
		userRepo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

		crypt := mock.NewMockCryptographyInterface(ctrl)
		crypt.EXPECT().Encrypt(gomock.Any()).Return(&token, nil).AnyTimes()

		useCase := uc.NewDefaulAuthenticateUseCase(userRepo, hasher, crypt)

		input := uc.AuthenticateUseCaseInput{
			Email:    "test@mail.com",
			Password: "invalid_password",
		}

		result, err := useCase.Execute(input)

		require.Equal(t, true, len(err) > 0)
		require.Nil(t, result)
	})

	t.Run("should not authenticate when password is incorrect", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		token := "token_test"
		userRepo := mock.NewMockUserRepositoryInterface(ctrl)
		userRepo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(false).AnyTimes()

		crypt := mock.NewMockCryptographyInterface(ctrl)
		crypt.EXPECT().Encrypt(gomock.Any()).Return(&token, nil).AnyTimes()

		useCase := uc.NewDefaulAuthenticateUseCase(userRepo, hasher, crypt)

		input := uc.AuthenticateUseCaseInput{
			Email:    "test@mail.com",
			Password: "Test#Incorrect",
		}

		result, err := useCase.Execute(input)

		require.Equal(t, true, len(err) > 0)
		require.Nil(t, result)
	})

	t.Run("should not authenticate when occurs errors in encrypt", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		userRepo := mock.NewMockUserRepositoryInterface(ctrl)
		userRepo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

		crypt := mock.NewMockCryptographyInterface(ctrl)
		crypt.EXPECT().Encrypt(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()

		useCase := uc.NewDefaulAuthenticateUseCase(userRepo, hasher, crypt)

		input := uc.AuthenticateUseCaseInput{
			Email:    "test@mail.com",
			Password: "Test#Incorrect",
		}

		result, err := useCase.Execute(input)

		require.Equal(t, true, len(err) > 0)
		require.Nil(t, result)
	})
}

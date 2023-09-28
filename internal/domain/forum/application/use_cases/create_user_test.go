package use_cases_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/intwone/ddd-golang/internal/constants"
	uc "github.com/intwone/ddd-golang/internal/domain/forum/application/use_cases"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create an user", func(t *testing.T) {
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).AnyTimes()

		fakeHashedPassword := "$2a$10$5ow21ZF9sH40Ka7LfhxukOZZrqmEqwaknsIJPolWD948qj.3cugem"
		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Hash(gomock.Any()).Return(&fakeHashedPassword, nil).AnyTimes()

		useCase := uc.NewDefaultCreateUserUseCase(repo, hasher)

		input := uc.CreateUserUseCaseInput{
			Name:     "Test Name",
			Email:    "test@mail.com",
			Password: "Test@123",
			Role:     "student",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not create an user when repo throw error", func(t *testing.T) {
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).Return(errors.New("any")).AnyTimes()

		fakeHashedPassword := "$2a$10$5ow21ZF9sH40Ka7LfhxukOZZrqmEqwaknsIJPolWD948qj.3cugem"
		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Hash(gomock.Any()).Return(&fakeHashedPassword, nil).AnyTimes()

		useCase := uc.NewDefaultCreateUserUseCase(repo, hasher)

		input := uc.CreateUserUseCaseInput{
			Name:     "Test Name",
			Email:    "test@mail.com",
			Password: "Test@123",
			Role:     "student",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})

	t.Run("should not create an user when email is invalid", func(t *testing.T) {
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).AnyTimes()

		fakeHashedPassword := "$2a$10$5ow21ZF9sH40Ka7LfhxukOZZrqmEqwaknsIJPolWD948qj.3cugem"
		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Hash(gomock.Any()).Return(&fakeHashedPassword, nil).AnyTimes()

		useCase := uc.NewDefaultCreateUserUseCase(repo, hasher)

		input := uc.CreateUserUseCaseInput{
			Name:     "Test Name",
			Email:    "invalid_mail.com",
			Password: "Test@123",
			Role:     "student",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, constants.InvalidEmailError, err.Error())
	})

	t.Run("should not create an user when password is invalid", func(t *testing.T) {
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).AnyTimes()

		fakeHashedPassword := "$2a$10$5ow21ZF9sH40Ka7LfhxukOZZrqmEqwaknsIJPolWD948qj.3cugem"
		hasher := mock.NewMockHasherInterface(ctrl)
		hasher.EXPECT().Hash(gomock.Any()).Return(&fakeHashedPassword, nil).AnyTimes()

		useCase := uc.NewDefaultCreateUserUseCase(repo, hasher)

		input := uc.CreateUserUseCaseInput{
			Name:     "Test Name",
			Email:    "valid@mail.com",
			Password: "test123",
			Role:     "student",
		}

		_, err := useCase.Execute(input)

		errors := [...]string{
			constants.NotContainMinimumCaracteresPasswordError,
			constants.NotContainUpperCaseCharacterePasswordError,
			constants.NotContainSpecialCharacterePasswordError,
		}

		e := errors[:]
		messageErr := strings.Join(e, ",")

		require.NotNil(t, err)
		require.Equal(t, messageErr, err.Error())
	})
}

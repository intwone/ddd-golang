package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
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
		useCase := uc.NewDefaultCreateUserUseCase(repo)

		input := uc.CreateUserUseCaseInput{
			Name: "Test Name",
			Role: "student",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not create an user when repo throw error", func(t *testing.T) {
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).Return(errors.New("any")).AnyTimes()
		useCase := uc.NewDefaultCreateUserUseCase(repo)

		input := uc.CreateUserUseCaseInput{
			Name: "Test Name",
			Role: "student",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}

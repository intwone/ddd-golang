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

func TestGetUserByEmailUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should get user by email", func(t *testing.T) {
		user, _ := enterprise.NewUser("Test Name", "test@mail.com", "Test@123", "student")
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(user, nil).AnyTimes()

		useCase := uc.NewDefaulGetUserByEmailUseCase(repo)

		input := uc.GetUserByEmailUseCaseInput{
			Email: "test@mail.com",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not get user by email", func(t *testing.T) {
		repo := mock.NewMockUserRepositoryInterface(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(nil, errors.New("any")).AnyTimes()

		useCase := uc.NewDefaulGetUserByEmailUseCase(repo)

		input := uc.GetUserByEmailUseCaseInput{
			Email: "test@mail.com",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}

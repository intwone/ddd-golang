package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/notification/application/use_cases"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestSendNotificationUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create a notification", func(t *testing.T) {
		repo := mock.NewMockNotificationsRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).AnyTimes()
		useCase := uc.NewDefaultSendNotificationUseCase(repo)

		input := uc.SendNotificationUseCaseInput{
			RecipientID: "1",
			Title:       "Title Test",
			Content:     "Content test",
		}

		_, err := useCase.Execute(input)

		require.Nil(t, err)
	})

	t.Run("should not create a notification", func(t *testing.T) {
		repo := mock.NewMockNotificationsRepositoryInterface(ctrl)
		repo.EXPECT().Create(gomock.Any()).Return(errors.New("any")).AnyTimes()
		useCase := uc.NewDefaultSendNotificationUseCase(repo)

		input := uc.SendNotificationUseCaseInput{
			RecipientID: "1",
			Title:       "Title Test",
			Content:     "Content test",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
	})
}

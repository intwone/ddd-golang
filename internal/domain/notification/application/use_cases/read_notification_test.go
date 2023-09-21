package use_cases_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/ddd-golang/internal/domain/notification/application/use_cases"
	"github.com/intwone/ddd-golang/internal/domain/notification/enterprise"
	mock "github.com/intwone/ddd-golang/internal/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestReadNotificationUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should read a notification", func(t *testing.T) {
		notification := enterprise.NewNotification("Title Test", "Content test", "1")
		repo := mock.NewMockNotificationsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*notification, nil).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultReadNotificationUseCase(repo)

		input := uc.ReadNotificationUseCaseInput{
			RecipientID:    "1",
			NotificationID: "1",
		}

		result, err := useCase.Execute(input)

		require.Nil(t, err)
		require.NotNil(t, result.GetReadAt())
	})

	t.Run("should not read a notification when not found notification", func(t *testing.T) {
		notification := enterprise.Notification{}
		repo := mock.NewMockNotificationsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(notification, errors.New("any")).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultReadNotificationUseCase(repo)

		input := uc.ReadNotificationUseCaseInput{
			RecipientID:    "1",
			NotificationID: "1",
		}

		result, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Nil(t, result.GetReadAt())
	})

	t.Run("should not read a notification from another user", func(t *testing.T) {
		notification := enterprise.NewNotification("Title Test", "Content test", "2", enterprise.NotificationOptionalParams{ID: "1"})
		repo := mock.NewMockNotificationsRepositoryInterface(ctrl)
		repo.EXPECT().GetByID(gomock.Any()).Return(*notification, nil).AnyTimes()
		repo.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		useCase := uc.NewDefaultReadNotificationUseCase(repo)

		input := uc.ReadNotificationUseCaseInput{
			RecipientID:    "1",
			NotificationID: "1",
		}

		_, err := useCase.Execute(input)

		require.NotNil(t, err)
		require.Equal(t, "not allowed", err.Error())
	})
}

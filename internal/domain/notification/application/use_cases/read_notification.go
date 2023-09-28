package use_cases

import (
	"errors"

	"github.com/intwone/ddd-golang/internal/constants"
	"github.com/intwone/ddd-golang/internal/domain/notification/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/notification/enterprise"
)

type ReadNotificationUseCaseInput struct {
	NotificationID string
	RecipientID    string
}

type ReadNotificationUseCaseInterface interface {
	Execute(input ReadNotificationUseCaseInput) (enterprise.Notification, error)
}

type DefaultReadNotificationUseCase struct {
	NotificationRepository repositories.NotificationsRepositoryInterface
}

func NewDefaultReadNotificationUseCase(notificationRepository repositories.NotificationsRepositoryInterface) *DefaultReadNotificationUseCase {
	return &DefaultReadNotificationUseCase{
		NotificationRepository: notificationRepository,
	}
}

func (uc *DefaultReadNotificationUseCase) Execute(input ReadNotificationUseCaseInput) (enterprise.Notification, error) {
	notification, getByIdErr := uc.NotificationRepository.GetByID(input.NotificationID)

	if getByIdErr != nil {
		return enterprise.Notification{}, getByIdErr
	}

	if notification.GetRecipientID() != input.RecipientID {
		return enterprise.Notification{}, errors.New(constants.NotAllowedError)
	}

	notification.Read()

	createErr := uc.NotificationRepository.Save(&notification)

	if createErr != nil {
		return enterprise.Notification{}, createErr
	}

	return notification, nil
}

package use_cases

import (
	"github.com/intwone/ddd-golang/internal/domain/notification/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/notification/enterprise"
)

type SendNotificationUseCaseInput struct {
	RecipientID string
	Title       string
	Content     string
}

type SendNotificationUseCaseInterface interface {
	Execute(input SendNotificationUseCaseInput) (enterprise.Notification, error)
}

type DefaultSendNotificationUseCase struct {
	NotificationRepository repositories.NotificationsRepositoryInterface
}

func NewDefaultSendNotificationUseCase(notificationRepository repositories.NotificationsRepositoryInterface) *DefaultSendNotificationUseCase {
	return &DefaultSendNotificationUseCase{
		NotificationRepository: notificationRepository,
	}
}

func (uc *DefaultSendNotificationUseCase) Execute(input SendNotificationUseCaseInput) (enterprise.Notification, error) {
	newNotification := enterprise.NewNotification(input.Title, input.Content, input.RecipientID)

	err := uc.NotificationRepository.Create(newNotification)

	if err != nil {
		return enterprise.Notification{}, err
	}

	return enterprise.Notification{}, nil
}

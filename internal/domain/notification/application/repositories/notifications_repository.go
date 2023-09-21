package repositories

import "github.com/intwone/ddd-golang/internal/domain/notification/enterprise"

type NotificationsRepositoryInterface interface {
	GetByID(id string) (enterprise.Notification, error)
	Create(notification *enterprise.Notification) error
	Save(answer *enterprise.Notification) error
}

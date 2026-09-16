package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
)

type NotificationService interface {
	CreateNotification(n *entities.Notification) error

	UpdateNotification(n *entities.Notification) error

	DeleteNotification(id uint) error

	GetNotificationByID(id uint) (*entities.Notification, error)

	GetFilteredNotifications(filters map[string]string) ([]entities.Notification, error)

	ListenOrderCreated()
}

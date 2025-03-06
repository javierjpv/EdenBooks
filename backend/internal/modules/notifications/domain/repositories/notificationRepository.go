package repositories

import "github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"

type NotificationRepository interface{
	
	CreateNotification(notification *entities.Notification)error

	UpdateNotification(notification *entities.Notification)error

	DeleteNotification(id uint)error

	GetNotificationByID(id uint)(*entities.Notification,error)
	
	GetFilteredNotifications(filters map[string]string) ([]entities.Notification, error)

}
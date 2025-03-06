package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/notifications/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
)

type NotificationService interface{
	
	CreateNotification(notification dto.NotificationDTO)error

	UpdateNotification(id uint, notification dto.NotificationDTO)error

	DeleteNotification(id uint)error

	GetNotificationByID(id uint)(*entities.Notification,error)

	GetFilteredNotifications(filters map[string]string) ([]entities.Notification, error)
	
	ListenOrderCreated() 

}
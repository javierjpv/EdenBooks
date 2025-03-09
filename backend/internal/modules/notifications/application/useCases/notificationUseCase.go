package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/notifications/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type NotificationUseCase struct {
	service services.NotificationService
}

func NewNotificationUseCase(service services.NotificationService) *NotificationUseCase {
	return &NotificationUseCase{service: service}
}

func (u *NotificationUseCase) CreateNotification(n dto.NotificationRequest) error {
	if n.Content == "" {
		return ErrMissingFields
	}
	if n.UserID == 0 {
		return ErrInvalid
	}
	return u.service.CreateNotification(n)
}

func (u *NotificationUseCase) UpdateNotification(id uint, n dto.NotificationRequest) error {
	if n.Content == "" {
		return ErrMissingFields
	}
	if n.UserID == 0 || id == 0 {
		return ErrInvalid
	}
	return u.service.UpdateNotification(id, n)
}

func (u *NotificationUseCase) DeleteNotification(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteNotification(id)
}

func (u *NotificationUseCase) GetNotificationByID(id uint) (*dto.NotificationResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	notification, err := u.service.GetNotificationByID(id)
	if err != nil {
		return nil, err
	}
	notificationResponse := dto.NewNotificationResponse(notification.ID,notification.CreatedAt,notification.UpdatedAt,notification.Content,notification.Seen,notification.UserID)
	return notificationResponse, nil
}
func (u *NotificationUseCase) GetFilteredNotifications(filters map[string]string) ([]dto.NotificationResponse, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"created_at": true, "updated_at": true, "name": true, "contact": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	// return u.service.GetFilteredNotificationes(filters)
	notificationes, err := u.service.GetFilteredNotifications(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Notification a NotificationResponse
	var notificationResponses []dto.NotificationResponse
	for _, notification := range notificationes {
		notificationResponses = append(notificationResponses, *dto.NewNotificationResponse(
			notification.ID,notification.CreatedAt,notification.UpdatedAt,notification.Content,notification.Seen,notification.UserID))
	}

	return notificationResponses, nil
}

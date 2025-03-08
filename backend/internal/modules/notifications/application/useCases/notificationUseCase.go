package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/notifications/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
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

func (u *NotificationUseCase) GetNotificationByID(id uint) (*entities.Notification, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	return u.service.GetNotificationByID(id)
}
func (u *NotificationUseCase) GetFilteredNotifications(filters map[string]string) ([]entities.Notification, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"created_at": true, "content": true, "seen": true, "user_id": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetFilteredNotifications(filters)
}

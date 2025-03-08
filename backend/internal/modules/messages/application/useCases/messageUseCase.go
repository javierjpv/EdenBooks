package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type MessageUseCase struct {
	service services.MessageService
}

func NewMessageUseCase(service services.MessageService) *MessageUseCase {
	return &MessageUseCase{service: service}
}

func (u *MessageUseCase) CreateMessage(m dto.MessageRequest) (*entities.Message, error) {
	if m.Content == "" || m.Status == "" {
		return nil, ErrMissingFields
	}
	if m.ChatID == 0 || m.ReceiverID == 0 || m.SenderID == 0 {
		return nil, ErrInvalid
	}
	return u.service.CreateMessage(m)
}

func (u *MessageUseCase) UpdateMessage(id uint, m dto.MessageRequest) error {
	if m.Content == "" || m.Status == "" {
		return ErrMissingFields
	}
	if m.ChatID == 0 || m.ReceiverID == 0 || m.SenderID == 0 {
		return ErrInvalid
	}
	return u.service.UpdateMessage(id, m)
}

func (u *MessageUseCase) DeleteMessage(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteMessage(id)
}

func (u *MessageUseCase) GetMessageByID(id uint) (*entities.Message, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	return u.service.GetMessageByID(id)
}

func (u *MessageUseCase) GetFilteredMessages(filters map[string]string) ([]entities.Message, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"content": true, "user_id": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetFilteredMessages(filters)
}

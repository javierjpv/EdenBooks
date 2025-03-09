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

func (u *MessageUseCase) GetMessageByID(id uint) (*dto.MessageResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	message, err := u.service.GetMessageByID(id)
	if err != nil {
		return nil, err
	}
	messageResponse := dto.NewMessageResponse(message.ID, message.CreatedAt, message.UpdatedAt, message.Content, message.Seen, message.Status, message.ChatID, message.SenderID, message.ReceiverID)
	return messageResponse, nil
}
func (u *MessageUseCase) GetFilteredMessages(filters map[string]string) ([]dto.MessageResponse, error) {
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

	messagees, err := u.service.GetFilteredMessages(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Message a MessageResponse
	var messageResponses []dto.MessageResponse
	for _, message := range messagees {
		messageResponses = append(messageResponses, *dto.NewMessageResponse(
			message.ID, message.CreatedAt, message.UpdatedAt, message.Content, message.Seen, message.Status, message.ChatID, message.SenderID, message.ReceiverID))
	}

	return messageResponses, nil
}

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
)

type MessageService interface {
	CreateMessage(message dto.MessageRequest) (*entities.Message, error)

	UpdateMessage(id uint, message dto.MessageRequest) error

	DeleteMessage(id uint) error

	GetMessageByID(id uint) (*entities.Message, error)

	GetFilteredMessages(filters map[string]string) ([]entities.Message, error)
}

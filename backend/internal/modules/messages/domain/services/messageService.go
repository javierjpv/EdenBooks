package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
)

type MessageService interface{
	CreateMessage(message dto.MessageDTO)(*entities.Message,error)

	UpdateMessage(id uint, message dto.MessageDTO)error

	DeleteMessage(id uint)error

	GetMessageByID(id uint)(*entities.Message,error)
	
	GetFilteredMessages(filters map[string]string) ([]entities.Message, error)
}
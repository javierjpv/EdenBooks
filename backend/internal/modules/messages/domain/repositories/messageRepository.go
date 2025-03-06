package repositories

import "github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"

type MessageRepository interface{
	CreateMessage(message *entities.Message)(*entities.Message,error)

	UpdateMessage(message *entities.Message)error

	DeleteMessage(id uint)error

	GetMessageByID(id uint)(*entities.Message,error)
	
	GetFilteredMessages(filters map[string]string) ([]entities.Message, error) 
}
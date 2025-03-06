package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/chats/application/dto"
)

type ChatService interface{
	CreateChat(chat dto.ChatDTO)(*entities.Chat,error)

    CreateChatWithUsers(userIDs []uint) (*entities.Chat,error)
	
	UpdateChat(id uint, chat dto.ChatDTO)error

	DeleteChat(id uint)error

	GetChatByID(id uint)(*entities.Chat,error)
	
	GetFilteredChats(filters map[string]string) ([]entities.Chat, error)
}
package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
)

type ChatService interface {
	CreateChat(chat *entities.Chat) (*entities.Chat, error)

	CreateChatWithUsers(userIDs []uint) (*entities.Chat, error)

	UpdateChat( chat *entities.Chat) error

	DeleteChat(id uint) error

	GetChatByID(id uint) (*entities.Chat, error)

	GetFilteredChats(filters map[string]string) ([]entities.Chat, error)
}

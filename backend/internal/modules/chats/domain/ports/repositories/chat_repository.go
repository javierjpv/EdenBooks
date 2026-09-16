package repositories

import "github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"

type ChatRepository interface{
	CreateChat(chat *entities.Chat)(*entities.Chat,error)
    
	CreateChatWithUsers(chat *entities.Chat, userIDs []uint) (*entities.Chat,error)
	
	UpdateChat(chat *entities.Chat)error

	DeleteChat(id uint)error

	GetChatByID(id uint)(*entities.Chat,error)
	
	GetFilteredChats(filters map[string]string) ([]entities.Chat, error) 
}
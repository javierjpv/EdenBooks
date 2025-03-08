package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/chats/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/repositories"
)

type ChatService struct {
	repo repositories.ChatRepository
}

func NewChatService(repo repositories.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) CreateChat(c dto.ChatRequest) (*entities.Chat, error) {
	chat := entities.NewChat()
	return s.repo.CreateChat(chat)
}

func (s *ChatService) CreateChatWithUsers(userIDs []uint) (*entities.Chat, error) {
	chat := entities.NewChat()
	return s.repo.CreateChatWithUsers(chat, userIDs)
}

func (s *ChatService) UpdateChat(id uint, t dto.ChatRequest) error {
	chat, err := s.repo.GetChatByID(id)
	if err != nil {
		return err
	}

	return s.repo.UpdateChat(chat)
}

func (s *ChatService) DeleteChat(id uint) error {
	if _, err := s.repo.GetChatByID(id); err != nil {
		return err
	}
	return s.repo.DeleteChat(id)
}

func (s *ChatService) GetChatByID(id uint) (*entities.Chat, error) {
	return s.repo.GetChatByID(id)
}

func (s *ChatService) GetFilteredChats(filters map[string]string) ([]entities.Chat, error) {
	return s.repo.GetFilteredChats(filters)
}

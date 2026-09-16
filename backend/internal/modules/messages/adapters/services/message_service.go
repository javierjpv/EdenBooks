package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/ports/repositories"
)

type MessageService struct {
	repo repositories.MessageRepository
}

func NewMessageService(repo repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message *entities.Message) (*entities.Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) UpdateMessage(m *entities.Message) error {
	return s.repo.UpdateMessage(m)
}

func (s *MessageService) DeleteMessage(id uint) error {
	return s.repo.DeleteMessage(id)
}

func (s *MessageService) GetMessageByID(id uint) (*entities.Message, error) {
	return s.repo.GetMessageByID(id)
}

func (s *MessageService) GetFilteredMessages(filters map[string]string) ([]entities.Message, error) {
	return s.repo.GetFilteredMessages(filters)
}

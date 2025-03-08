package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/repositories"
)

type MessageService struct {
	repo repositories.MessageRepository
}

func NewMessageService(repo repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(m dto.MessageRequest) (*entities.Message, error) {
	message := entities.NewMessage(m.Content, m.Seen, m.Status, m.ChatID, m.SenderID, m.ReceiverID)
	return s.repo.CreateMessage(message)
}

func (s *MessageService) UpdateMessage(id uint, m dto.MessageRequest) error {
	message, err := s.repo.GetMessageByID(id)
	if err != nil {
		return err
	}
	message.Content = m.Content
	message.Seen = m.Seen
	message.Status = m.Status
	message.ChatID = m.ChatID
	message.SenderID = m.SenderID
	message.ReceiverID = m.ReceiverID
	return s.repo.UpdateMessage(message)
}

func (s *MessageService) DeleteMessage(id uint) error {
	if _, err := s.repo.GetMessageByID(id); err != nil {
		return err
	}
	return s.repo.DeleteMessage(id)
}

func (s *MessageService) GetMessageByID(id uint) (*entities.Message, error) {
	return s.repo.GetMessageByID(id)
}

func (s *MessageService) GetFilteredMessages(filters map[string]string) ([]entities.Message, error) {
	return s.repo.GetFilteredMessages(filters)
}

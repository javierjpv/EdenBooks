package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/chats/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/services"
	messageDto "github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	userDto "github.com/javierjpv/edenBooks/internal/modules/users/application/dto"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type ChatUseCase struct {
	service services.ChatService
}

func NewChatUseCase(service services.ChatService) *ChatUseCase {
	return &ChatUseCase{service: service}
}

func (u *ChatUseCase) CreateChat(c dto.ChatRequest) (*entities.Chat, error) {
	return u.service.CreateChat(c)
}
func (u *ChatUseCase) CreateChatWithUsers(userIDs []uint) (*entities.Chat, error) {
	return u.service.CreateChatWithUsers(userIDs)
}

func (u *ChatUseCase) UpdateChat(id uint, c dto.ChatRequest) error {
	return u.service.UpdateChat(id, c)
}

func (u *ChatUseCase) DeleteChat(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteChat(id)
}

func (u *ChatUseCase) GetChatByID(id uint) (*dto.ChatResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	chat, err := u.service.GetChatByID(id)
	if err != nil {
		return nil, err
	}
	// Convertir cada usuario en UserResponse
	var users []userDto.UserResponse
	for _, user := range chat.Users {
		users = append(users, userDto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
		})
	}
	// Convertir cada mensaje en MessageResponse (si es necesario)
	var messages []messageDto.MessageResponse
	for _, message := range chat.Messages {
		messages = append(messages, messageDto.MessageResponse{
			ID:        message.ID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
			SenderID:  message.SenderID,
		})
	}
	// Convertir cada Chat a ChatResponse
	chatResponse := dto.NewChatResponse(chat.ID, chat.CreatedAt, chat.UpdatedAt, users, messages)
	return chatResponse, nil
}
func (u *ChatUseCase) GetFilteredChats(filters map[string]string) ([]dto.ChatResponse, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"created_at": true, "updated_at": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	chats, err := u.service.GetFilteredChats(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Chat a ChatResponse
	var chatResponses []dto.ChatResponse
	for _, chat := range chats {
		// Convertir usuarios en UserResponse
		var users []userDto.UserResponse
		for _, user := range chat.Users {
			users = append(users, userDto.UserResponse{
				ID:    user.ID,
				Email: user.Email,
			})
		}

		// Convertir mensajes en MessageResponse
		var messages []messageDto.MessageResponse
		for _, message := range chat.Messages {
			messages = append(messages, messageDto.MessageResponse{
				ID:        message.ID,
				Content:   message.Content,
				CreatedAt: message.CreatedAt,
				SenderID:  message.SenderID,
			})
		}

		chatResponses = append(chatResponses, *dto.NewChatResponse(chat.ID, chat.CreatedAt, chat.UpdatedAt, users, messages))
	}

	return chatResponses, nil
}

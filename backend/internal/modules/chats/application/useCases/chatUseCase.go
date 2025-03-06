package usecases

import (
	"fmt"
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/chats/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/services"
)

var (
    ErrMissingFields = fmt.Errorf("all fields are required")
    ErrInvalid     = fmt.Errorf("invalid ID")
)
type ChatUseCase struct {
	service services.ChatService
}


func NewChatUseCase(service services.ChatService)*ChatUseCase{
	return &ChatUseCase{service: service}
}

func (u *ChatUseCase)CreateChat(c dto.ChatDTO)(*entities.Chat,error){
	return u.service.CreateChat(c)
}
func (u *ChatUseCase)CreateChatWithUsers(userIDs []uint) (*entities.Chat,error){
	return u.service.CreateChatWithUsers(userIDs)
}

func (u *ChatUseCase)UpdateChat(id uint,c dto.ChatDTO)error{
	return u.service.UpdateChat(id,c)
}

func (u *ChatUseCase)DeleteChat(id uint)error{
	if  id==0  {
		return ErrInvalid
	}
	return u.service.DeleteChat(id)
}

func (u *ChatUseCase)GetChatByID(id uint)(*entities.Chat,error){
	if  id==0  {
		return nil,ErrInvalid
	}
	return u.service.GetChatByID(id)
}

func (u *ChatUseCase) GetFilteredChats(filters map[string]string) ([]entities.Chat, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"creted_at": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetFilteredChats(filters)
}

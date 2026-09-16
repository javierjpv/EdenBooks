package dto

import (
	"time"

	messageDto "github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	userDto "github.com/javierjpv/edenBooks/internal/modules/users/application/dto"
)
 type ChatResponse struct {
	 ID         uint
	 CreatedAt  time.Time
	 UpdatedAt  time.Time
     Users []userDto.UserResponse
     Messages []messageDto.MessageResponse
 }



 func NewChatResponse(ID uint, CreatedAt, UpdatedAt time.Time, Users []userDto.UserResponse, Messages []messageDto.MessageResponse) *ChatResponse {
	return &ChatResponse{
		ID:        ID,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
		Users:     Users,
		Messages:  Messages,
	}
}
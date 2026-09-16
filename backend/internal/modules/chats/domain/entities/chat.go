package entities

import (
	"time"

	messageEntities "github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
)

type Chat struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []userEntities.User       `gorm:"many2many:user_chats;"`
	Messages  []messageEntities.Message `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` ///Relacion 1:N

}

func NewChat() *Chat {
	return &Chat{}
}

package entities

import (

	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	messageEntities "github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"

	"gorm.io/gorm"
)


type Chat struct{
	gorm.Model
	
	Users []userEntities.User `gorm:"many2many:user_chats;"` 
	Messages []messageEntities.Message `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` ///Relacion 1:N
	
}





func NewChat()*Chat{
	return &Chat{}
}


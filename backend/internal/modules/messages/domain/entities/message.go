package entities

import (
	"gorm.io/gorm"
)

type Message struct{
	gorm.Model
	Content string      `gorm:"not null"`
	Seen bool
	Status string
	ChatID uint         `gorm:"not null"`//Relacion 1:N
	SenderID   uint     `gorm:"not null"` // Relación 1:N con Usuario (quién lo envía)
    ReceiverID uint     `gorm:"not null"`// Relación 1:N con Usuario (quién lo recibe)
}
func NewMessage(content string,seen bool,status string,chatID uint,senderID uint, receiverID uint)*Message{
	return &Message{Content: content,Seen: seen,Status: status,ChatID: chatID,SenderID: senderID,ReceiverID: receiverID}
}
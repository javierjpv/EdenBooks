package dto

import "time"

type MessageResponse struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Content    string
	Seen       bool
	Status     string
	ChatID     uint
	SenderID   uint
	ReceiverID uint
}

func NewMessageResponse(ID uint, CreatedAt time.Time, UpdatedAt time.Time, content string, seen bool, status string, chatID uint, senderID uint, receiverID uint) *MessageResponse {
	return &MessageResponse{ID: ID,CreatedAt: CreatedAt,UpdatedAt: UpdatedAt,Content: content, Seen: seen, Status: status, ChatID: chatID, SenderID: senderID, ReceiverID: receiverID}
}

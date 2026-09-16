package entities

import (
	"time"
)

type Notification struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	Seen      bool
	UserID    uint // Relacion 1:N
}

func NewNotification(content string, seen bool, userID uint) *Notification {
	return &Notification{Content: content, Seen: seen, UserID: userID}
}

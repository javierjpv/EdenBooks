package dto

import "time"

type NotificationResponse struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	Seen      bool
	UserID    uint
}

func NewNotificationResponse(ID uint, CreatedAt time.Time, UpdatedAt time.Time, Content string, Seen bool, UserID uint) *NotificationResponse {
	return &NotificationResponse{ID: ID,CreatedAt: CreatedAt,UpdatedAt: UpdatedAt,Content: Content,Seen: Seen,UserID: UserID}
}

package dto

type NotificationRequest struct {
	Content string `json:"content"`
	Seen    bool   `json:"seen"`
	UserID  uint   `json:"user_id"`
}

package entities

import "gorm.io/gorm"

type Notification struct{
	gorm.Model
	Content string
	Seen bool
	UserID uint // Relacion 1:N
}

func NewNotification(content string,seen bool,userID uint)*Notification{
return &Notification{Content: content,Seen: seen,UserID: userID}
}
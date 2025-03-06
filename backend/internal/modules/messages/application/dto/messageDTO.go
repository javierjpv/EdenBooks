package dto


type MessageDTO struct{
	Content string   `json:"Content"`
	Seen bool        `json:"Seen"`
	Status string    `json:"Status"`
	ChatID uint      `json:"ChatID"` 
	SenderID   uint  `json:"SenderID"`
    ReceiverID uint  `json:"ReceiverID"` 
}
func NewMessageDTO(content string,seen bool,status string,chatID uint,senderID uint, receiverID uint)*MessageDTO{
	return &MessageDTO{Content: content,Seen: seen,Status: status,ChatID: chatID,SenderID: senderID,ReceiverID: receiverID}
}
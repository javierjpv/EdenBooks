package repositories

import (
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	"gorm.io/gorm"
)

type MessageRepository struct{
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB)*MessageRepository{
	return &MessageRepository{db: db}
}

func (r *MessageRepository)CreateMessage(message *entities.Message)(*entities.Message,error){
	if err:=r.db.Create(message).Error; err!=nil {
		return nil, err
	}
	return message,nil

}
func (r *MessageRepository)UpdateMessage(message *entities.Message) error{
	err:=r.db.Save(message).Error
	if  err!=nil {
		return err
	}
	return nil
}

func (r *MessageRepository)DeleteMessage(id uint) error{
	err:=r.db.Delete(&entities.Message{},id).Error
	if err!=nil {
		return err
	}
	return nil
}

func (r *MessageRepository)GetMessageByID(id uint) (*entities.Message,error){
	var message entities.Message 
	err:=r.db.First(&message,id).Error
	if err!=nil {
		return nil,err
	}
	return &message,nil
}
func (r *MessageRepository) GetFilteredMessages(filters map[string]string) ([]entities.Message, error) {
	var messages []entities.Message
	query := r.db

	// Aplicar filtros dinámicos
	for key, value := range filters {
		switch key {
		case "content":
			query = query.Where("content = ?", value)
		case "user_id":
			query = query.Where("user_id >= ?", value)
		}
	}

	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := filters["order"]
		query = query.Order(sortBy + " " + order)
	}

	// Aplicar paginación si está presente
	limitInt := 50 // Límite por defecto
	if limit, exists := filters["limit"]; exists {
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = parsedLimit
		}
	}
	query = query.Limit(limitInt)

	if page, exists := filters["page"]; exists {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			query = query.Offset((pageInt - 1) * limitInt)
		}
	}

	// Ejecutar la consulta
	if err := query.Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

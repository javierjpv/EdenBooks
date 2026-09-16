package repositories

import (
	"log"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/chats/domain/entities"
	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) CreateChat(chat *entities.Chat) (*entities.Chat, error) {
	if err := r.db.Create(chat).Error; err != nil {
		return nil, err
	}
	return chat, nil

}

func (r *ChatRepository) CreateChatWithUsers(chat *entities.Chat, userIDs []uint) (*entities.Chat, error) {
	log.Printf("🔹 Iniciando la creación del chat con usuarios: %v\n", userIDs)
	var existingChat entities.Chat

	err := r.db.
		Joins("JOIN user_chats ON user_chats.chat_id = chats.id").
		Where("user_chats.user_id IN (?)", userIDs).
		Group("chats.id").
		Having("COUNT(DISTINCT user_chats.user_id) = ?", len(userIDs)).
		First(&existingChat).Error

	// Si se encontró un chat, retornarlo
	if err == nil {
		log.Printf("Chat existente encontrado con ID: %d\n", existingChat.ID)
		return &existingChat, nil
	}

	// Si ocurre otro error que no sea "registro no encontrado", retornar el error
	if err != gorm.ErrRecordNotFound {
		log.Printf("Error al buscar el chat: %v\n", err)
		return nil, err
	}
	log.Printf("🔹 No se encontró chat existente, creando uno nuevo...")

	var users []userEntities.User
	if err := r.db.Find(&users, userIDs).Error; err != nil {
		return nil, err
	}

	// Asociar los usuarios al chat
	chat.Users = users
	// Crear el chat
	if err := r.db.Create(&chat).Error; err != nil {
		return nil, err
	}
	log.Printf("Chat creado con ID: %d\n", chat.ID)

	if err := r.db.Preload("Users").First(&chat, chat.ID).Error; err != nil {
		return nil, err
	}

	return chat, nil

}

func (r *ChatRepository) UpdateChat(chat *entities.Chat) error {
	result := r.db.Model(&entities.Chat{}).Where("id = ?", chat.ID).Updates(chat)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ChatRepository) DeleteChat(id uint) error {
	result := r.db.Delete(&entities.Chat{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *ChatRepository) GetChatByID(id uint) (*entities.Chat, error) {
	var chat entities.Chat
	err := r.db.Preload("Users").Preload("Messages").First(&chat, id).Error
	if err != nil {
		return nil, err
	}
	return &chat, nil
}
func (r *ChatRepository) GetFilteredChats(filters map[string]string) ([]entities.Chat, error) {
	var chats []entities.Chat
	query := r.db
	for key, value := range filters {
		switch key {
		case "user_id":
			query = query.Joins("JOIN user_chats ON user_chats.chat_id = chats.id").
				Where("user_chats.user_id = ?", value)
		}
	}

	query = query.Preload("Users")
	query = query.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Where("id IN (?)", r.db.Table("messages").Select("MAX(id)").Group("chat_id")).Order("created_at desc")
	})
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
	if err := query.Find(&chats).Error; err != nil {
		return nil, err
	}

	return chats, nil
}

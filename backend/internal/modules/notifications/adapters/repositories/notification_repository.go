package repositories

import (
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(notification *entities.Notification) error {
	if err := r.db.Create(notification).Error; err != nil {
		return err
	}
	return nil

}
func (r *NotificationRepository) UpdateNotification(n *entities.Notification) error {
    result := r.db.Model(&entities.Notification{}).Where("id = ?", n.ID).Updates(n)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}

func (r *NotificationRepository) DeleteNotification(id uint) error {
    result := r.db.Delete(&entities.Notification{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}
func (r *NotificationRepository) GetNotificationByID(id uint) (*entities.Notification, error) {
	var notification entities.Notification
	err := r.db.First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *NotificationRepository) GetFilteredNotifications(filters map[string]string) ([]entities.Notification, error) {
	var notifications []entities.Notification
	query := r.db

	// Aplicar filtros dinámicos
	for key, value := range filters {
		switch key {
		case "start_date":
			query = query.Where("created_at >= ?", value)
		case "end_date":
			query = query.Where("created_at <= ?", value)
		case "content":
			query = query.Where("content LIKE ?", "%"+value+"%")
		case "seen":
			query = query.Where("seen = ?", value)
		case "user_id":
			query = query.Where("user_id = ?", value)
		}
	}

	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := filters["order"]
		query = query.Order(sortBy + " " + order)
	}

	// Aplicar paginación si está presente
	if page, exists := filters["page"]; exists {
		limit := filters["limit"]
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		query = query.Offset((pageInt - 1) * limitInt).Limit(limitInt)
	}

	// Ejecutar la consulta
	if err := query.Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}

package repositories

import (
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category *entities.Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) UpdateCategory(category *entities.Category) error {
	result := r.db.Model(&entities.Category{}).Where("id = ?", category.ID).Updates(category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(id uint) error {
	result := r.db.Delete(&entities.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*entities.Category, error) {
	var category entities.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAllCategories() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

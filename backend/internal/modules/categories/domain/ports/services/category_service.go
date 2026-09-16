package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
)

type CategoryService interface {
	CreateCategory(c *entities.Category) error

	UpdateCategory(c *entities.Category) error

	DeleteCategory(id uint) error

	GetCategoryByID(id uint) (*entities.Category, error)

	GetAllCategories() ([]entities.Category, error)
}

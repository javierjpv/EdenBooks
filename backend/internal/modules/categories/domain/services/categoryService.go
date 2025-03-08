package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/categories/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
)

type CategoryService interface {
	CreateCategory(dto.CategoryRequest) error

	UpdateCategory(id uint, c dto.CategoryRequest) error

	DeleteCategory(id uint) error

	GetCategoryByID(id uint) (*entities.Category, error)

	GetAllCategories() ([]entities.Category, error)
}

package repositories

import "github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"

type CategoryRepository interface{

	CreateCategory(category *entities.Category) error

	UpdateCategory(category *entities.Category) error

	DeleteCategory(id uint) error

	GetCategoryByID(id uint) (*entities.Category,error)
    
	GetAllCategories()([]entities.Category,error)
}
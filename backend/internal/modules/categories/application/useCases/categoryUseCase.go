package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/categories/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type CategoryUseCase struct {
	service services.CategoryService
}

func NewCategoryUseCase(service services.CategoryService) *CategoryUseCase {
	return &CategoryUseCase{service: service}
}

func (u *CategoryUseCase) CreateCategory(category dto.CategoryRequest) error {
	if category.Name == "" || category.Description == "" {
		return ErrMissingFields
	}
	return u.service.CreateCategory(category)
}

func (u *CategoryUseCase) UpdateCategory(id uint, c dto.CategoryRequest) error {
	if id == 0 {
		return ErrInvalid
	}
	if c.Name == "" || c.Description == "" {
		return ErrMissingFields
	}
	return u.service.UpdateCategory(id, c)
}

func (u *CategoryUseCase) DeleteCategory(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteCategory(id)
}

func (u *CategoryUseCase) GetCategoryByID(id uint) (*dto.CategoryResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	category, err := u.service.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.NewCategoryResponse(category.ID, category.Name,category.Description)
	return categoryResponse, nil
}
func (u *CategoryUseCase) GetAllCategories() ([]dto.CategoryResponse, error) {

	categories, err := u.service.GetAllCategories()
	if err != nil {
		return nil, err
	}

	// Convertir cada Category a CategoryResponse
	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, *dto.NewCategoryResponse(
			category.ID,category.Name,category.Description))
	}

	return categoryResponses, nil
}

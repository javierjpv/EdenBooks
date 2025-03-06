package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/categories/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/services"
)

var (
    ErrMissingFields = fmt.Errorf("all fields are required")
    ErrInvalid     = fmt.Errorf("invalid ID")
)
type CategoryUseCase struct{
	service services.CategoryService
}

func NewCategoryUseCase(service services.CategoryService)*CategoryUseCase{
	return &CategoryUseCase{service: service}
}

func (u *CategoryUseCase)CreateCategory(category dto.CategoryDTO)error{
	if category.Name =="" || category.Description =="" {
		return ErrMissingFields
	}
	return u.service.CreateCategory(category)
}

func (u *CategoryUseCase)UpdateCategory(id uint, c dto.CategoryDTO)error{
	if id==0 {
		return ErrInvalid
	}
	if c.Name =="" || c.Description =="" {
		return ErrMissingFields
	}
	return u.service.UpdateCategory(id,c)
}

func (u *CategoryUseCase)DeleteCategory(id uint)error{
	if id==0 {
		return ErrInvalid
	}
	return u.service.DeleteCategory(id)
}

func (u *CategoryUseCase)GetCategoryByID(id uint)(*entities.Category,error){
	if id==0 {
		return nil,ErrInvalid
	}
	return u.service.GetCategoryByID(id)
}

func (u *CategoryUseCase)GetAllCategories()([]entities.Category,error){
	return u.service.GetAllCategories()
}
// GetAllCategories()([]entities.Category)
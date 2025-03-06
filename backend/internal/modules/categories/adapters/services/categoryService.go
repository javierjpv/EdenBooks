package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/categories/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/repositories"
)

type CategoryService struct{
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository)*CategoryService{
	return &CategoryService{repo: repo}
}
func (s *CategoryService)CreateCategory(c dto.CategoryDTO) error{
    category:=entities.NewCategory(c.Name,c.Description)
	return s.repo.CreateCategory(category)
}

func (s *CategoryService)UpdateCategory(id uint, c dto.CategoryDTO) error{
    category,err:=s.repo.GetCategoryByID(id)
	category.Name=c.Name
	category.Description=c.Description
	if err!=nil {
		return err
	}
	return s.repo.UpdateCategory(category)
}

func (s *CategoryService)DeleteCategory(id uint) error{
	if _,err:=s.repo.GetCategoryByID(id);err!=nil {
		return err
	}
    return s.repo.DeleteCategory(id)
}

func (s *CategoryService)GetCategoryByID(id uint) (*entities.Category,error){
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService)GetAllCategories()([]entities.Category,error){
	return s.repo.GetAllCategories()
}
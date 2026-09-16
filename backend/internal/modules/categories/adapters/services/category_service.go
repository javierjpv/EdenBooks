package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/categories/domain/ports/repositories"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}
func (s *CategoryService) CreateCategory(c *entities.Category) error {
	return s.repo.CreateCategory(c)
}

func (s *CategoryService) UpdateCategory(c *entities.Category) error {
	return s.repo.UpdateCategory(c)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}

func (s *CategoryService) GetCategoryByID(id uint) (*entities.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) GetAllCategories() ([]entities.Category, error) {
	return s.repo.GetAllCategories()
}

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/ports/repositories"
	eventBusService "github.com/javierjpv/edenBooks/internal/shared/domain/services"
)

type ProductService struct {
	repo            repositories.ProductRepository
	eventBusService eventBusService.EventBus
}

func NewProductService(repo repositories.ProductRepository, eventBusService eventBusService.EventBus) *ProductService {
	return &ProductService{repo: repo, eventBusService: eventBusService}
}

func (s *ProductService) CreateProduct(p *entities.Product) error {
	return s.repo.CreateProduct(p)
}

func (s *ProductService) UpdateProduct(p *entities.Product) error {
	return s.repo.UpdateProduct(p)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}
func (s *ProductService) GetProductByID(id uint) (*entities.Product, error) {
	return s.repo.GetProductByID(id)
}
func (s *ProductService) GetProductByIDWithFavorite(id uint, userID uint) (*entities.Product, bool, error) {
	return s.repo.GetProductByIDWithFavorite(id, userID)
}
func (s *ProductService) AddToFavorites(userID uint, productID uint) error {
	if _, err := s.repo.GetProductByID(productID); err != nil {
		return err
	}
	return s.repo.AddToFavorites(userID, productID)
}
func (s *ProductService) RemoveFromFavorites(userID uint, productID uint) error {
	if _, err := s.repo.GetProductByID(productID); err != nil {
		return err
	}
	return s.repo.RemoveFromFavorites(userID, productID)
}

func (s *ProductService) GetFavorites(userID uint) ([]entities.Product, error) {
	return s.repo.GetFavorites(userID)
}

func (s *ProductService) GetFilteredProducts(filters map[string]string) ([]entities.Product, error) {
	return s.repo.GetFilteredProducts(filters)
}
func (s *ProductService) GetProductsWithFavorites(userID uint, filters map[string]string) ([]entities.Product, []bool, error) {
	return s.repo.GetProductsWithFavorites(userID, filters)
}

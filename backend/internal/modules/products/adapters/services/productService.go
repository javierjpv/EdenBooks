package services

import (

	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/repositories"
	eventBusService "github.com/javierjpv/edenBooks/internal/shared/domain/services"
)

type ProductService struct {
	repo            repositories.ProductRepository
	eventBusService eventBusService.EventBus
}

func NewProductService(repo repositories.ProductRepository, eventBusService eventBusService.EventBus) *ProductService {
	return &ProductService{repo: repo, eventBusService: eventBusService}
}

func (s *ProductService) CreateProduct(p dto.ProductRequest) error {
	product := entities.NewProduct(p.Name, p.Description, p.Price, p.CategoryID, p.UserID, p.ImageURL) //se creara un user sin order id ya q esta se llevara a cabo mas adelante
	return s.repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(id uint, p dto.ProductRequest) error {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return err
	}
	product.Name = p.Name
	product.Description = p.Description
	product.Price = p.Price
	product.CategoryID = p.CategoryID
	product.UserID = p.UserID
	product.ImageURL = p.ImageURL
	product.Sold = p.Sold
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	if _, err := s.repo.GetProductByID(id); err != nil {
		return err
	}
	return s.repo.DeleteProduct(id)
}
func (s *ProductService) GetProductByID(id uint) (*entities.Product, error) {
	return s.repo.GetProductByID(id)
}
func (s *ProductService) GetProductByIDWithFavorite(id uint, userID uint) (*dto.ProductResponse, error) {
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
func (s *ProductService) GetProductsWithFavorites(userID uint, filters map[string]string) ([]dto.ProductResponse, error) {
	return s.repo.GetProductsWithFavorites(userID, filters)
}

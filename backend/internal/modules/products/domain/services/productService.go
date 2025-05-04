package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
)

type ProductService interface {
	CreateProduct(dto.ProductRequest) error
	UpdateProduct(id uint, product dto.ProductRequest) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*entities.Product, error)
	GetProductByIDWithFavorite(id uint, userID uint) (*dto.ProductResponse, error)
	AddToFavorites(userID uint, productID uint) error
	RemoveFromFavorites(userID uint, productID uint) error
	GetFavorites(userID uint) ([]entities.Product, error)
	GetFilteredProducts(filters map[string]string) ([]entities.Product, error)
	GetProductsWithFavorites(userID uint, filters map[string]string) ([]dto.ProductResponse, error)
}

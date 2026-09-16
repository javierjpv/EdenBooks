package services

import "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
type ProductService interface {
	CreateProduct(p *entities.Product) error
	UpdateProduct(p *entities.Product) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*entities.Product, error)
	GetProductByIDWithFavorite(id uint, userID uint) (*entities.Product, bool, error)
	AddToFavorites(userID uint, productID uint) error
	RemoveFromFavorites(userID uint, productID uint) error
	GetFavorites(userID uint) ([]entities.Product, error)
	GetFilteredProducts(filters map[string]string) ([]entities.Product, error)
	GetProductsWithFavorites(userID uint, filters map[string]string) ([]entities.Product, []bool, error)
}

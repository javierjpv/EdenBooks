package repositories

import "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"


type ProductRepository interface{
	CreateProduct(product *entities.Product)error
	UpdateProduct(product *entities.Product)error
	DeleteProduct(id uint)error
	GetProductByID(id uint)(*entities.Product,error)
	AddToFavorites(userID uint,productID uint)error
	RemoveFromFavorites(userID uint, productID uint) error
	GetFavorites(userID uint) ([]entities.Product, error)
	GetFilteredProducts(filters map[string]string) ([]entities.Product, error)
	GetProductsWithFavorites(userID uint, filters map[string]string) ([]entities.ProductWithFavoriteStatus, error)
}
package repositories

import "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"


type ProductRepository interface{
	CreateProduct(product *entities.Product)error
	UpdateProduct(product *entities.Product)error
	DeleteProduct(id uint)error
	GetProductByID(id uint)(*entities.Product,error)
	GetFilteredProducts(filters map[string]string) ([]entities.Product, error)
}
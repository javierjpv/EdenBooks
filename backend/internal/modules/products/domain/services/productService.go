package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
)

type ProductService interface{
	CreateProduct(dto.ProductDTO)error
	UpdateProduct(id uint, product dto.ProductDTO)error
	AddOrderIDToProducts(orderID uint,productsIDs []uint)error
	DeleteProduct(id uint)error
	GetProductByID(id uint)(*entities.Product,error)
	GetFilteredProducts(filters map[string]string) ([]entities.Product, error)
}
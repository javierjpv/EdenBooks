package dto

import "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"

type ProductResponse struct {
	entities.Product
	IsFavorite bool `json:"is_favorite"`
}

func NewProductResponse(product *entities.Product, isFavorite bool) *ProductResponse {
	if product == nil {
		return nil
	}
	return &ProductResponse{
		Product:    *product,
		IsFavorite: isFavorite,
	}
}

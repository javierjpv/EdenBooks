package dto

import "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"

type ProductResponse struct {
	entities.Product
	IsFavorite bool `json:"is_favorite"`
}

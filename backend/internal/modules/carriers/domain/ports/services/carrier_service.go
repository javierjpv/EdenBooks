package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"
)

type CarrierService interface {
	CreateCarrier(c *entities.Carrier) error

	UpdateCarrier(c *entities.Carrier) error

	DeleteCarrier(id uint) error

	GetCarrierByID(id uint) (*entities.Carrier, error)

	GetFilteredCarrieres(filters map[string]string) ([]entities.Carrier, error)
}

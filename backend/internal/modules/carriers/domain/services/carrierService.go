package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/carriers/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"
)

type CarrierService interface{
	
	CreateCarrier(carrier dto.CarrierDTO)error

	UpdateCarrier(id uint,a dto.CarrierDTO)error

	DeleteCarrier(id uint)error

	GetCarrierByID(id uint)(*entities.Carrier,error)

	GetFilteredCarrieres(filters map[string]string) ([]entities.Carrier, error)

}
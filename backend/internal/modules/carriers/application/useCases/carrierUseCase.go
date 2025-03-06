package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/carriers/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/services"
)

var (
    ErrMissingFields = fmt.Errorf("all fields are required")
    ErrInvalidID     = fmt.Errorf("invalid Carrier ID")
)

type CarrierUseCase struct{
	service services.CarrierService
}

func NewCarrierUseCase(service services.CarrierService)*CarrierUseCase{
	return &CarrierUseCase{service: service}
}

func (u *CarrierUseCase) CreateCarrier(carrier dto.CarrierDTO)error{
	if carrier.Contact == "" || carrier.Name == "" {
        return ErrMissingFields
    }
	return u.service.CreateCarrier(carrier)
}

func (u *CarrierUseCase) UpdateCarrier(id uint, carrier dto.CarrierDTO) error {
	if id==0 {
		return ErrInvalidID
	}
	if carrier.Contact == "" || carrier.Name == ""{
        return ErrMissingFields
    }
	return u.service.UpdateCarrier(id,carrier)
}

func (u *CarrierUseCase) DeleteCarrier(id uint) error {
	if id==0 {
		return ErrInvalidID
	}
	return u.service.DeleteCarrier(id)
}

func (u *CarrierUseCase) GetCarrierByID(id uint) (*entities.Carrier, error) {
	if id==0 {
		return nil,ErrInvalidID
	}
	return u.service.GetCarrierByID(id)
}
func (u *CarrierUseCase) GetFilteredCarrieres(filters map[string]string) ([]entities.Carrier, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"created_at": true, "updated_at": true, "name": true, "contact": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetFilteredCarrieres(filters)
}
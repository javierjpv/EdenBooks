package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/carriers/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalidID     = fmt.Errorf("invalid Carrier ID")
)

type CarrierUseCase struct {
	service services.CarrierService
}

func NewCarrierUseCase(service services.CarrierService) *CarrierUseCase {
	return &CarrierUseCase{service: service}
}

func (u *CarrierUseCase) CreateCarrier(carrier dto.CarrierRequest) error {
	if carrier.Contact == "" || carrier.Name == "" {
		return ErrMissingFields
	}
	return u.service.CreateCarrier(carrier)
}

func (u *CarrierUseCase) UpdateCarrier(id uint, carrier dto.CarrierRequest) error {
	if id == 0 {
		return ErrInvalidID
	}
	if carrier.Contact == "" || carrier.Name == "" {
		return ErrMissingFields
	}
	return u.service.UpdateCarrier(id, carrier)
}

func (u *CarrierUseCase) DeleteCarrier(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	return u.service.DeleteCarrier(id)
}

func (u *CarrierUseCase) GetCarrierByID(id uint) (*dto.CarrierResponse, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	carrier, err := u.service.GetCarrierByID(id)
	if err != nil {
		return nil, err
	}
	carrierResponse := dto.NewCarrierResponse(carrier.ID, carrier.Name, carrier.Contact)
	return carrierResponse, nil
}
func (u *CarrierUseCase) GetFilteredCarrieres(filters map[string]string) ([]dto.CarrierResponse, error) {
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

	// return u.service.GetFilteredCarrieres(filters)
	carrieres, err := u.service.GetFilteredCarrieres(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Carrier a CarrierResponse
	var carrierResponses []dto.CarrierResponse
	for _, carrier := range carrieres {
		carrierResponses = append(carrierResponses, *dto.NewCarrierResponse(
			carrier.ID,carrier.Name,carrier.Contact))
	}

	return carrierResponses, nil
}

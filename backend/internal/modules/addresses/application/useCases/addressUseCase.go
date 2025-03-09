package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalidID     = fmt.Errorf("invalid Address ID")
)

type AddressUseCase struct {
	service services.AddressService
}

func NewAddressUseCase(service services.AddressService) *AddressUseCase {
	return &AddressUseCase{service: service}
}

func (u *AddressUseCase) CreateAddress(address dto.AddressRequest) (*dto.AddressResponse, error) {
	if address.City == "" || address.Province == "" || address.PostalCode == "" || address.Country == "" || address.Street == "" {
		return nil, ErrMissingFields
	}
	if address.Number <= 0 {
		return nil, fmt.Errorf("adrees number can not be <= 0")
	}
	createdAddress, err := u.service.CreateAddress(address)
	if err != nil {
		return nil, err
	}
	addressResponse := dto.NewAddressResponse(createdAddress.ID, createdAddress.City, createdAddress.Province, createdAddress.PostalCode, createdAddress.Country, createdAddress.Street, createdAddress.Number)
	return addressResponse, nil
}

func (u *AddressUseCase) UpdateAddress(id uint, address dto.AddressRequest) error {
	if id == 0 {
		return ErrInvalidID
	}
	if address.City == "" || address.Province == "" || address.PostalCode == "" || address.Country == "" || address.Street == "" {
		return ErrMissingFields
	}
	if address.Number <= 0 {
		return fmt.Errorf("adrees number can not be <= 0")
	}
	return u.service.UpdateAddress(id, address)
}

func (u *AddressUseCase) DeleteAddress(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	return u.service.DeleteAddress(id)
}

func (u *AddressUseCase) GetAddressByID(id uint) (*dto.AddressResponse, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}
	address, err := u.service.GetAddressByID(id)
	if err != nil {
		return nil, err
	}
	addressResponse := dto.NewAddressResponse(address.ID, address.City, address.Province, address.PostalCode, address.Country, address.Street, address.Number)
	return addressResponse, nil
}
func (u *AddressUseCase) GetFilteredAddresses(filters map[string]string) ([]dto.AddressResponse, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"created_at": true, "updated_at": true, "city": true, "province": true, "postal_code": true, "country": true, "street": true, "number": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	addresses, err := u.service.GetFilteredAddresses(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Address a AddressResponse
	var addressResponses []dto.AddressResponse
	for _, address := range addresses {
		addressResponses = append(addressResponses, *dto.NewAddressResponse(
			address.ID, address.City, address.Province, address.PostalCode, address.Country, address.Street, address.Number))
	}

	return addressResponses, nil
}

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
)

type AddressService interface {
	CreateAddress(address dto.AddressRequest) (*entities.Address, error)

	UpdateAddress(id uint, a dto.AddressRequest) error

	DeleteAddress(id uint) error

	GetAddressByID(id uint) (*entities.Address, error)

	CheckExistingAdress(a dto.AddressRequest) (bool, error)

	GetFilteredAddresses(filters map[string]string) ([]entities.Address, error)
}

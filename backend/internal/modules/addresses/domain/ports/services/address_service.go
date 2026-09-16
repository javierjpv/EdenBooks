package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
)

type AddressService interface {
	CreateAddress(a *entities.Address) (*entities.Address, error)

	UpdateAddress(a *entities.Address) error

	DeleteAddress(id uint) error

	GetAddressByID(id uint) (*entities.Address, error)

	CheckExistingAdress(a entities.Address) (bool, error)

	GetFilteredAddresses(filters map[string]string) ([]entities.Address, error)
}

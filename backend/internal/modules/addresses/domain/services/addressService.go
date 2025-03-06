package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
)

type AddressService interface{
	
	CreateAddress(address dto.AddressDTO)(*entities.Address, error) 

	UpdateAddress(id uint,a dto.AddressDTO)error

	DeleteAddress(id uint)error

	GetAddressByID(id uint)(*entities.Address,error)

	CheckExistingAdress(a dto.AddressDTO)(bool,error)

	GetFilteredAddresses(filters map[string]string) ([]entities.Address, error)

}
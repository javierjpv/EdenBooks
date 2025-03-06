package repositories

import "github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"

type AddressRepository interface{
	
	CreateAddress(address *entities.Address) (*entities.Address, error) 

	UpdateAddress(address *entities.Address) error

	DeleteAddress(id uint) error

	GetAddressByID(id uint) (*entities.Address,error)

	CheckExistingAddress(address *entities.Address) (*entities.Address, bool, error)

    GetFilteredAddresses(filters map[string]string) ([]entities.Address,error)

}



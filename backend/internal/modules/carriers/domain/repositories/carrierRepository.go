package repositories

import "github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"

type CarrierRepository interface{
	
	CreateCarrier(Carrier *entities.Carrier) error

	UpdateCarrier(Carrier *entities.Carrier) error

	DeleteCarrier(id uint) error

	GetCarrierByID(id uint) (*entities.Carrier,error)

    GetFilteredCarrieres(filters map[string]string) ([]entities.Carrier,error)

}



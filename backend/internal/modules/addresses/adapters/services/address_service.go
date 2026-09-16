package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/ports/repositories"
)

type AddressService struct {
	repo repositories.AddressRepository //Como se epuede observar se importa la interfaz   //Mas adelante al constructor le pasaras un objeto de tipo AddressRepository
	//Las ventajas de usar una interfaz son desacoplamiento, mejor mocking,testing, inversion de dependencias,flexibilidad
}

func NewAddressService(repo repositories.AddressRepository) *AddressService { //Al declarar el argumento como interfaz el metodo sera capaz de aceptar cualquier struct que contenga los metodos con sus correspondiente comportamiento
	return &AddressService{repo: repo}
}

func (s *AddressService) CreateAddress(a *entities.Address) (*entities.Address, error) {
	return s.repo.CreateAddress(a)

}

func (s *AddressService) UpdateAddress(a *entities.Address) error {
	return s.repo.UpdateAddress(a)

}

func (s *AddressService) DeleteAddress(id uint) error {
	return s.repo.DeleteAddress(id)
}

func (s *AddressService) GetAddressByID(id uint) (*entities.Address, error) {
	return s.repo.GetAddressByID(id)

}
func (s *AddressService) CheckExistingAdress(a entities.Address) (bool, error) {
	return false, nil

}
func (s *AddressService) GetFilteredAddresses(filters map[string]string) ([]entities.Address, error) {
	return s.repo.GetFilteredAddresses(filters)
}
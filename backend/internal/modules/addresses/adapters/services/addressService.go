package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/repositories"
)


type AddressService struct{
	repo repositories.AddressRepository //Como se epuede observar se importa la interfaz   //Mas adelante al constructor le pasaras un objeto de tipo AddressRepository
	//Las ventajas de usar una interfaz son desacoplamiento, mejor mocking,testing, inversion de dependencias,flexibilidad 
}

func NewAddressService(repo repositories.AddressRepository)*AddressService{ //Al declarar el argumento como interfaz el metodo sera capaz de aceptar cualquier struct que contenga los metodos con sus correspondiente comportamiento
	return &AddressService{repo: repo}
}

func (s *AddressService) CreateAddress(a dto.AddressDTO)(*entities.Address, error) {
	address:=entities.NewAddress(a.City, a.Province, a.PostalCode, a.Country)
	return s.repo.CreateAddress(address) //no hace falta & ya que el metodo NewAdrees ya devuelve un puntero

}
  
func (s *AddressService) UpdateAddress(id uint,a dto.AddressDTO)error{
	address,err:=s.repo.GetAddressByID(id)
	if err!=nil {
		return err
	}
	address.City=a.City
	address.Province=a.Province
	address.PostalCode=a.PostalCode
	address.Country=a.Country
	return s.repo.UpdateAddress(address)

}

func (s *AddressService) DeleteAddress(id uint)error{
	_,err:=s.repo.GetAddressByID(id)
	if err!=nil {
		return err
	}
	return s.repo.DeleteAddress(id)
}

func (s *AddressService) GetAddressByID(id uint)(*entities.Address,error){
	return s.repo.GetAddressByID(id)

}
func (s *AddressService) CheckExistingAdress(a dto.AddressDTO)(bool,error){
	return false,nil

}
func (s *AddressService) GetFilteredAddresses(filters map[string]string) ([]entities.Address, error) {
	return s.repo.GetFilteredAddresses(filters)
}

// type Repository[T any] interface {
//     Create(entity *T) error
//     Update(entity *T) error
//     Delete(id uint) error
//     GetByID(id uint) (*T, error)
// }

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/carriers/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/repositories"
)

type CarrierService struct {
	repo repositories.CarrierRepository //Como se epuede observar se importa la interfaz   //Mas adelante al constructor le pasaras un objeto de tipo CarrierRepository
	//Las ventajas de usar una interfaz son desacoplamiento, mejor mocking,testing, inversion de dependencias,flexibilidad
}

func NewCarrierService(repo repositories.CarrierRepository) *CarrierService { //Al declarar el argumento como interfaz el metodo sera capaz de aceptar cualquier struct que contenga los metodos con sus correspondiente comportamiento
	return &CarrierService{repo: repo}
}

func (s *CarrierService) CreateCarrier(c dto.CarrierRequest) error {
	carrier := entities.NewCarrier(c.Name, c.Contact)
	return s.repo.CreateCarrier(carrier) //no hace falta & ya que el metodo NewAdrees ya devuelve un puntero

}

func (s *CarrierService) UpdateCarrier(id uint, c dto.CarrierRequest) error {
	carrier, err := s.repo.GetCarrierByID(id)
	if err != nil {
		return err
	}
	carrier.Name = c.Name
	carrier.Contact = c.Contact
	return s.repo.UpdateCarrier(carrier)

}

func (s *CarrierService) DeleteCarrier(id uint) error {
	_, err := s.repo.GetCarrierByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteCarrier(id)
}

func (s *CarrierService) GetCarrierByID(id uint) (*entities.Carrier, error) {
	return s.repo.GetCarrierByID(id)

}

func (s *CarrierService) GetFilteredCarrieres(filters map[string]string) ([]entities.Carrier, error) {
	return s.repo.GetFilteredCarrieres(filters)
}

// type Repository[T any] interface {
//     Create(entity *T) error
//     Update(entity *T) error
//     Delete(id uint) error
//     GetByID(id uint) (*T, error)
// }

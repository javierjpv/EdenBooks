package services

import (
	"errors"
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/ports/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
func (s *UserService) Register(u entities.User) error {
	user := entities.NewUser(u.Email, u.Password)
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return errors.New("usuario no creado")
	}
	return nil
}
func (s *UserService) Login(email string, password string) (*entities.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("contraseña incorrecta")
	}

	return user, nil

}
func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetUserByID(id)
}
func (s *UserService) UpdateUser(u *entities.User) error {
	return s.userRepo.UpdateUser(u)
}

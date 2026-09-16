package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
)

type UserService interface {
	Register(user entities.User) error
	Login(email string, password string) (*entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	UpdateUser(u *entities.User) error
}

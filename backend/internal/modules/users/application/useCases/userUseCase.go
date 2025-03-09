package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/users/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid")
)

type UserUseCase struct {
	service services.UserService
}

func NewUserUseCase(service services.UserService) *UserUseCase {
	return &UserUseCase{service: service}
}

func (u *UserUseCase) Register(user dto.UserRequest) error {

	if user.Email == "" || user.Password == "" {
		return ErrMissingFields
	}
	return u.service.Register(user)
}

func (u *UserUseCase) Login(email string, password string) (*entities.User, error) {

	if email == "" || password == "" {
		return nil, ErrInvalid
	}
	return u.service.Login(email, password)

}
func (u *UserUseCase) GetUserByID(id uint) (*dto.UserResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}

	user,err:=u.service.GetUserByID(id)
	if err!=nil {
		return nil,err
	}
	userResponse:=dto.UserResponse{ID:user.ID,Email: user.Email }
	return &userResponse, nil
}

func (u *UserUseCase) UpdateUser(id uint, user dto.UserRequest) error {
	return u.service.UpdateUser(id, user)
}

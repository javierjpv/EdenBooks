package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/users/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
)



type UserService interface{

	Register(dto.UserDTO)error
	Login(email string, password string)(*entities.User,error)
	GetUserByID(id uint)(*entities.User,error)
	UpdateUser(id uint, product dto.UserDTO)error

	
}
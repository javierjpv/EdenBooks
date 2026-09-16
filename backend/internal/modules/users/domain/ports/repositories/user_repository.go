package repositories

import (
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
)

type UserRepository interface{

	CreateUser(user *entities.User)error
	FindByEmail(email string)(*entities.User,error)
	GetUserByID(id uint)(*entities.User,error)
	UpdateUser(user *entities.User)error

}
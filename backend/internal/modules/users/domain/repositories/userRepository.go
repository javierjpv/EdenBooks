package repositories

import (
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
)

type UserRepository interface{

	CreateUser(*entities.User)error
	FindByEmail(email string)(*entities.User,error)
	GetUserByID(id uint)(*entities.User,error)
	UpdateUser(product *entities.User)error

}
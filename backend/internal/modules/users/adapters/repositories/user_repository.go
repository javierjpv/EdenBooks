package repositories

import (
	"github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//aqui hashear contraseñas antes de guardarlas

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	// Hashear la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) GetUserByID(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) UpdateUser(u *entities.User) error {
	// Si la contraseña viene rellena, la hasheamos antes de actualizar
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	result := r.db.Model(&entities.User{}).Where("id = ?", u.ID).Updates(u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

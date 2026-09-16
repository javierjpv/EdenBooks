package entities

import (
	"time"
	productEntities "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
)

type Category struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Products    []productEntities.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N
}

func NewCategory(name string, description string) *Category {
	return &Category{Name: name, Description: description}

}

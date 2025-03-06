package entities

import (
	productEntities "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"gorm.io/gorm"
)


type Category struct{
	gorm.Model
	Name string
	Description string
	Products []productEntities.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N
}

func NewCategory(name string,description string)*Category{
return &Category{Name: name,Description: description}

}
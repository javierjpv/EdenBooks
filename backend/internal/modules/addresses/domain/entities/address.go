package entities

import (
	"time"

	orderEntities "github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
)

type Address struct {
		ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	City       string
	Province   string
	PostalCode string
	Country    string
	Street     string
	Number     int
	Users      []userEntities.User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N  una direccion puede tener muchos usuarios
	Orders     []orderEntities.Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N  una direccion puede tener muchos pedidos
}

func NewAddress(city string, province string, postalCode string, country string) *Address {
	return &Address{City: city, Province: province, PostalCode: postalCode, Country: country}
}

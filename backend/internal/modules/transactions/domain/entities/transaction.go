package entities

import (
	orderEntities "github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	"gorm.io/gorm"
)

type Transaction struct{
	gorm.Model
	PaymentMethod string //lo mejor sera crear una tabla con los metodos de pago
    Total float64
	Order orderEntities.Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relación 1:1
}

func NewTransaction(paymentMethod string,total float64)*Transaction{
	return &Transaction{PaymentMethod: paymentMethod,Total: total}
}

package entities

import (
	// productEntities "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"gorm.io/gorm"
)


type Order struct{
	gorm.Model
	State string
	UserID uint //// Relacion 1:N
	AddressID uint
	CarrierID      uint
	TransactionID uint
	// No hare el preload asi que lo comento por si en el futuro se hace
	// Products []productEntities.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`//Relacion 1:N un pedido tendra muchos productos pero un producto solo podra estar en un pedido(esto debido a que son productos unicos y no es un eccomerce donde un producto puede pertenecer a muchos pedidos(en esos casos la relacion seria N:M))
}

func NewOrder(state string,userID uint,addressID uint,carrierID uint,TransactionID uint )*Order{
	return &Order{State: state, UserID: userID, AddressID:addressID,CarrierID: carrierID,TransactionID: TransactionID}
}
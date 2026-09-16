package entities

import "time"

type Order struct {
	ID            uint `gorm:"primaryKey"`
	State         string
	UserID        uint
	AddressID     uint
	CarrierID     uint
	TransactionID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewOrder(state string, userID, addressID, carrierID, transactionID uint) *Order {
	return &Order{
		State:         state,
		UserID:        userID,
		AddressID:     addressID,
		CarrierID:     carrierID,
		TransactionID: transactionID,
	}
}

// import (
// 	// productEntities "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
// 	"gorm.io/gorm"
// )

// type Order struct{
// 	gorm.Model
// 	State string
// 	UserID uint //// Relacion 1:N
// 	AddressID uint
// 	CarrierID      uint
// 	TransactionID uint
// }

// func NewOrder(state string,userID uint,addressID uint,carrierID uint,TransactionID uint )*Order{
// 	return &Order{State: state, UserID: userID, AddressID:addressID,CarrierID: carrierID,TransactionID: TransactionID}
// }

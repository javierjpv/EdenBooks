package dto

import "time"


type OrderResponse struct {
	ID            uint
	CreatedAt time.Time
	UpdatedAt time.Time
	State         string 
	UserID        uint   
	AddressID     uint   
	CarrierID     uint   
	TransactionID uint   
}

func NewOrderResponse(id uint, createdAt, updatedAt time.Time, state string, userID, addressID, carrierID, transactionID uint) *OrderResponse {
	return &OrderResponse{
		ID:            id,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		State:         state,
		UserID:        userID,
		AddressID:     addressID,
		CarrierID:     carrierID,
		TransactionID: transactionID,
	}
}
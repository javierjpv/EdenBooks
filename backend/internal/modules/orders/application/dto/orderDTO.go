package dto

type OrderDTO struct {
	State         string  `json:"state"`
	UserID        uint    `json:"userID"`
	AddressID     uint    `json:"addressID"`
	CarrierID     uint    `json:"carrierID"`
	TransactionID uint    `json:"transactionID"`
}

func NewOrderDTO(state string,userID uint,addressID uint,carrierID uint,TransactionID uint )*OrderDTO{
	return &OrderDTO{State: state, UserID: userID, AddressID:addressID,CarrierID: carrierID,TransactionID: TransactionID}
}
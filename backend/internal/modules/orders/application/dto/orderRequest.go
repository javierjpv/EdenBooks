package dto

type OrderRequest struct {
	State         string `json:"state"`
	UserID        uint   `json:"userID"`
	AddressID     uint   `json:"addressID"`
	CarrierID     uint   `json:"carrierID"`
	OrderID     uint     `json:"orderID"`
	TransactionID uint   `json:"transactionID"`
}

func NewOrderRequest(state string, userID uint, addressID uint, carrierID uint, TransactionID uint) *OrderRequest {
	return &OrderRequest{State: state, UserID: userID, AddressID: addressID, CarrierID: carrierID, TransactionID: TransactionID}
}

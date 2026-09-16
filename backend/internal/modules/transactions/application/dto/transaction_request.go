package dto

type TransactionRequest struct {
	PaymentMethod string  `json:"payment_method"`
	Total         float64 `json:"total"`
}

func NewTransactionRequest(paymentMethod string, total float64) *TransactionRequest {
	return &TransactionRequest{PaymentMethod: paymentMethod, Total: total}
}

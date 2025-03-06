package dto

type TransactionDTO struct{
	PaymentMethod string `json:"payment_method"`
    Total float64        `json:"total"`
}

func NewTransactionDTO(paymentMethod string,total float64)*TransactionDTO{
	return &TransactionDTO{PaymentMethod: paymentMethod,Total: total}
}

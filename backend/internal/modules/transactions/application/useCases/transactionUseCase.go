package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/transactions/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type TransactionUseCase struct {
	service services.TransactionService
}

func NewTransactionUseCase(service services.TransactionService) *TransactionUseCase {
	return &TransactionUseCase{service: service}
}

func (u *TransactionUseCase) CreateTransaction(t dto.TransactionRequest) (*entities.Transaction, error) {
	if t.PaymentMethod == "" {
		return nil, ErrMissingFields
	}
	if t.Total < 0 {
		return nil, fmt.Errorf("total can not be < 0")
	}
	return u.service.CreateTransaction(t)
}

func (u *TransactionUseCase) UpdateTransaction(id uint, t dto.TransactionRequest) error {
	if t.PaymentMethod == "" {
		return ErrMissingFields
	}
	if t.Total < 0 {
		return fmt.Errorf("total can not be < 0")
	}
	if id == 0 {
		return ErrInvalid
	}
	return u.service.UpdateTransaction(id, t)
}

func (u *TransactionUseCase) DeleteTransaction(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteTransaction(id)
}

func (u *TransactionUseCase) GetTransactionByID(id uint) (*entities.Transaction, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	return u.service.GetTransactionByID(id)
}

func (u *TransactionUseCase) GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"total": true, "payment_method": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetFilteredTransactions(filters)
}

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/transactions/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
)

type TransactionService interface {
	CreateTransaction(transaction dto.TransactionRequest) (*entities.Transaction, error)

	UpdateTransaction(id uint, transaction dto.TransactionRequest) error

	DeleteTransaction(id uint) error

	GetTransactionByID(id uint) (*entities.Transaction, error)

	GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error)
}

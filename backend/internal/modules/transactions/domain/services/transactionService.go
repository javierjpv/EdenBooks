package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/application/dto"
)

type TransactionService interface{
	CreateTransaction(transaction dto.TransactionDTO)(*entities.Transaction,error)

	UpdateTransaction(id uint, transaction dto.TransactionDTO)error

	DeleteTransaction(id uint)error

	GetTransactionByID(id uint)(*entities.Transaction,error)
	
	GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error)
}
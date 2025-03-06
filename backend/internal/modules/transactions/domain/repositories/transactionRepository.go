package repositories

import "github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"

type TransactionRepository interface{
	CreateTransaction(transaction *entities.Transaction)(*entities.Transaction,error)

	UpdateTransaction(transaction *entities.Transaction)error

	DeleteTransaction(id uint)error

	GetTransactionByID(id uint)(*entities.Transaction,error)
	
	GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error) 
}
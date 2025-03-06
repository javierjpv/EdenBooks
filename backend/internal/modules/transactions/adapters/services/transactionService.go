package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/repositories"
)

type TransactionService struct{
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository)*TransactionService{
	return &TransactionService{repo: repo}
}

func(s * TransactionService)CreateTransaction(t dto.TransactionDTO)(*entities.Transaction,error){
	transaction:=entities.NewTransaction(t.PaymentMethod,t.Total)
    return s.repo.CreateTransaction(transaction)
}

func(s * TransactionService)UpdateTransaction(id uint, t dto.TransactionDTO)error{
	transaction,err:=s.repo.GetTransactionByID(id)
    if err!=nil{
		return err
	}
    transaction.PaymentMethod=t.PaymentMethod
	transaction.Total=t.Total
	return s.repo.UpdateTransaction(transaction)	
}

func(s * TransactionService)DeleteTransaction(id uint)error{
	if _,err:=s.repo.GetTransactionByID(id);err!=nil{
		return err
	}
	return s.repo.DeleteTransaction(id)
}

func(s * TransactionService)GetTransactionByID(id uint)(*entities.Transaction,error){
	return s.repo.GetTransactionByID(id)
}

func (s *TransactionService) GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error) {
	return s.repo.GetFilteredTransactions(filters)
}

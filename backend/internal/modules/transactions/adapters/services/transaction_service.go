package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/ports/repositories"
)

type TransactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(t *entities.Transaction) (*entities.Transaction, error) {
	return s.repo.CreateTransaction(t)
}

func (s *TransactionService) UpdateTransaction(t *entities.Transaction) error {
	return s.repo.UpdateTransaction(t)
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	if _, err := s.repo.GetTransactionByID(id); err != nil {
		return err
	}
	return s.repo.DeleteTransaction(id)
}

func (s *TransactionService) GetTransactionByID(id uint) (*entities.Transaction, error) {
	return s.repo.GetTransactionByID(id)
}

func (s *TransactionService) GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error) {
	return s.repo.GetFilteredTransactions(filters)
}

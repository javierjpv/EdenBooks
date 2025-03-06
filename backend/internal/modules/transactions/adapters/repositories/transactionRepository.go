package repositories

import (
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/transactions/domain/entities"
	"gorm.io/gorm"
)

type TransactionRepository struct{
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB)*TransactionRepository{
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository)CreateTransaction(transaction *entities.Transaction)(*entities.Transaction,error){
	if err:=r.db.Create(transaction).Error; err!=nil {
		return nil, err
	}
	return transaction,nil

}
func (r *TransactionRepository)UpdateTransaction(transaction *entities.Transaction) error{
	err:=r.db.Save(transaction).Error
	if  err!=nil {
		return err
	}
	return nil
}

func (r *TransactionRepository)DeleteTransaction(id uint) error{
	err:=r.db.Delete(&entities.Transaction{},id).Error
	if err!=nil {
		return err
	}
	return nil
}

func (r *TransactionRepository)GetTransactionByID(id uint) (*entities.Transaction,error){
	var transaction entities.Transaction 
	err:=r.db.First(&transaction,id).Error
	if err!=nil {
		return nil,err
	}
	return &transaction,nil
}
func (r *TransactionRepository) GetFilteredTransactions(filters map[string]string) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	query := r.db

	// Aplicar filtros dinámicos
	for key, value := range filters {
		switch key {
		case "payment_method":
			query = query.Where("payment_method = ?", value)
		case "min_total":
			query = query.Where("total >= ?", value)
		case "max_total":
			query = query.Where("total <= ?", value)
		}
	}

	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := filters["order"]
		query = query.Order(sortBy + " " + order)
	}

	// Aplicar paginación si está presente
	limitInt := 50 // Límite por defecto
	if limit, exists := filters["limit"]; exists {
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = parsedLimit
		}
	}
	query = query.Limit(limitInt)

	if page, exists := filters["page"]; exists {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			query = query.Offset((pageInt - 1) * limitInt)
		}
	}

	// Ejecutar la consulta
	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

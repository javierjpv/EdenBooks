package repositories

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository struct{
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB)*ProductRepository{
	return &ProductRepository{db: db}
}

func (r *ProductRepository)CreateProduct(product *entities.Product)error{
	if err:=r.db.Create(product).Error; err!=nil {
		return err
	}
	return nil

}
func (r *ProductRepository)UpdateProduct(product *entities.Product) error{
	err:=r.db.Save(product).Error
	if  err!=nil {
		return err
	}
	return nil
}

func (r *ProductRepository)DeleteProduct(id uint) error{
	err:=r.db.Delete(&entities.Product{},id).Error
	if err!=nil {
		return err
	}
	return nil
}

func (r *ProductRepository)GetProductByID(id uint) (*entities.Product,error){
    var product entities.Product
    err:=r.db.First(&product,id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) { // Check if it's a "record not found" error
            return nil, fmt.Errorf("product with ID %d not found", id) // User-friendly message
        }
        return nil, fmt.Errorf("error fetching product: %w", err) // Wrap other errors
    }
    return &product,nil
}
func (r *ProductRepository) GetFilteredProducts(filters map[string]string) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db

	for key, value := range filters {
		switch key {
		case "name":
			query = query.Where("name ILIKE ?", "%"+value+"%") 
		case "description":
			query = query.Where("description ILIKE ?", "%"+value+"%")
		case "min_price":
			query = query.Where("price >= ?", value)
		case "max_price":
			query = query.Where("price <= ?", value)
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "user_id":
			query = query.Where("user_id = ?", value)
		case "order_id":
			query = query.Where("order_id = ?", value)
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
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

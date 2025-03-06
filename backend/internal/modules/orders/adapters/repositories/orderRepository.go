package repositories

import (
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	"gorm.io/gorm"
)

type OrderRepository struct{
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB)*OrderRepository{
	return &OrderRepository{db: db}
}

func (r *OrderRepository)CreateOrder(order *entities.Order)(uint,error){
	if err:=r.db.Create(order).Error; err!=nil {
		return 0,err
	}
	return order.ID,nil

}
func (r *OrderRepository)UpdateOrder(order *entities.Order) error{
	err:=r.db.Save(order).Error
	if  err!=nil {
		return err
	}
	return nil
}

func (r *OrderRepository)DeleteOrder(id uint) error{
	err:=r.db.Delete(&entities.Order{},id).Error
	if err!=nil {
		return err
	}
	return nil
}

func (r *OrderRepository)GetOrderByID(id uint) (*entities.Order,error){
	var order entities.Order 
	err:=r.db.First(&order,id).Error
	if err!=nil {
		return nil,err
	}
	return &order,nil
}
func (r *OrderRepository) GetFilteredOrders(filters map[string]string) (*[]entities.Order, error) {
	var orders []entities.Order
	query := r.db

	// Aplicar filtros dinámicos
	for key, value := range filters {
		switch key {
		case "state":
			query = query.Where("state = ?", value)
		case "address_id":
			query = query.Where("address_id = ?", value)
		case "carrier_id":
			query = query.Where("carrier_id = ?", value)
		case "transaction_id":
			query = query.Where("transaction_id = ?", value)
		case "user_id":
			query = query.Where("user_id = ?", value)
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
	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

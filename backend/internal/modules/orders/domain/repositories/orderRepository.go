package repositories

import (
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
)

type OrderRepository interface{
	
	CreateOrder(order *entities.Order)(uint,error)

	UpdateOrder(order *entities.Order)error

	DeleteOrder(id uint)error

	GetOrderByID(id uint)(*entities.Order,error)
	
    GetFilteredOrders(filters map[string]string) (*[]entities.Order, error)
}
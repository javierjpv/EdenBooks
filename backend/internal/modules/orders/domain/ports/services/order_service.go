package services

import "github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"


type OrderService interface {
	CheckOrder(order *entities.Order, productIDs []uint) error

	AddOrderIDToProducts(orderID uint, productsIDs []uint) error

	CreateOrder(order *entities.Order, productIDs []uint) error

	UpdateOrder(order *entities.Order) error

	DeleteOrder(id uint) error

	GetOrderByID(id uint) (*entities.Order, error)

	GetFilteredOrders(filters map[string]string) (*[]entities.Order, error)

	ListenPaymentCreated()
}

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/orders/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
)

type OrderService interface {
	CheckOrder(o dto.OrderRequest, productsIDs []uint) error

	CreateOrder(o dto.OrderRequest, productsIDs []uint) error

	UpdateOrder(id uint, order dto.OrderRequest) error

	DeleteOrder(id uint) error

	GetOrderByID(id uint) (*entities.Order, error)

	GetFilteredOrders(filters map[string]string) (*[]entities.Order, error)

	ListenPaymentCreated()
}

package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/orders/application/dto"
)

type OrderService interface{
	CheckOrder(o dto.OrderDTO, productsIDs []uint) (error)

	CreateOrder(o dto.OrderDTO, productsIDs []uint)error

	UpdateOrder(id uint, order dto.OrderDTO)error
    
	DeleteOrder(id uint)error

	GetOrderByID(id uint)(*entities.Order,error)
    
	GetFilteredOrders(filters map[string]string) (*[]entities.Order, error)
	
	ListenPaymentCreated()
}
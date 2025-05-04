package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/orders/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type OrderUseCase struct {
	service services.OrderService
}

func NewOrderUseCase(service services.OrderService) *OrderUseCase {
	return &OrderUseCase{service: service}
}

func (u *OrderUseCase) CheckOrder(o dto.OrderRequest, productsIDs []uint) error {
	if o.State == "" {
		return ErrMissingFields
	}
	if o.AddressID == 0 || o.UserID == 0 {
		return ErrInvalid
	}
	for _, productID := range productsIDs {
		if productID == 0 {
			return ErrInvalid
		}
	}
	return u.service.CheckOrder(o, productsIDs)
}
func (u *OrderUseCase) CreateOrder(o dto.OrderRequest, productsIDs []uint) error {
	if o.State == "" {
		return ErrMissingFields
	}
	if o.AddressID == 0 || o.UserID == 0 || o.TransactionID == 0 {
		return ErrInvalid
	}
	for _, productID := range productsIDs {
		if productID == 0 {
			return ErrInvalid
		}
	}
	return u.service.CreateOrder(o, productsIDs)
}

func (u *OrderUseCase) UpdateOrder(id uint, o dto.OrderRequest) error {
	if o.State == "" {
		return ErrMissingFields
	}
	if o.AddressID == 0 || o.UserID == 0 || id == 0 || o.TransactionID == 0 {
		return ErrInvalid
	}
	return u.service.UpdateOrder(id, o)
}

func (u *OrderUseCase) DeleteOrder(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteOrder(id)
}

func (u *OrderUseCase) GetOrderByID(id uint) (*dto.OrderResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	order, err := u.service.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	orderResponse := dto.NewOrderResponse(order.ID, order.CreatedAt, order.UpdatedAt, order.State, order.UserID, order.AddressID, order.CarrierID, order.TransactionID)
	return orderResponse, nil
}
func (u *OrderUseCase) GetFilteredOrders(filters map[string]string) ([]dto.OrderResponse, error) {
	// Validar el orden si está presente
	if ord, exists := filters["order"]; exists {
		if ord != "asc" && ord != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"state": true, "carrier_id": true, "transaction_id": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	// return u.service.GetFilteredOrderes(filters)
	orders, err := u.service.GetFilteredOrders(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Order a OrderResponse
	var orderResponses []dto.OrderResponse
	for _, order := range *orders {
		orderResponses = append(orderResponses, *dto.NewOrderResponse(
			order.ID, order.CreatedAt, order.UpdatedAt, order.State,
			order.UserID, order.AddressID, order.CarrierID, order.TransactionID,
		))
	}

	return orderResponses, nil
}

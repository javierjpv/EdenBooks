package services

import (
	"fmt"
	"log"

	addressEntities "github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
	addressServices "github.com/javierjpv/edenBooks/internal/modules/addresses/domain/ports/services"
	carrierServices "github.com/javierjpv/edenBooks/internal/modules/carriers/domain/ports/services"
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/ports/repositories"
	productServices "github.com/javierjpv/edenBooks/internal/modules/products/domain/ports/services"
	userServices "github.com/javierjpv/edenBooks/internal/modules/users/domain/ports/services"
)

type OrderService struct {
	repo            repositories.OrderRepository
	productService  productServices.ProductService
	userService     userServices.UserService
	addressService  addressServices.AddressService
	carrierService  carrierServices.CarrierService
	eventBusService EventBus
}
type EventBus interface {
	Publish(topic string, data interface{}) error
	Subscribe(topic string, handler func(data interface{})) error
}

func NewOrderService(repo repositories.OrderRepository, productService productServices.ProductService, addressService addressServices.AddressService, carrierService carrierServices.CarrierService, userService userServices.UserService, eventBusService EventBus) *OrderService {
	return &OrderService{repo: repo, productService: productService, addressService: addressService, carrierService: carrierService, userService: userService, eventBusService: eventBusService}
}

func (s *OrderService) AddOrderIDToProducts(orderID uint, productsIDs []uint) error {
	order, err := s.repo.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("order does not exist: %w", err)
	}

	for _, id := range productsIDs {
		p, err := s.productService.GetProductByID(id)
		if err != nil {
			return fmt.Errorf("product %d does not exist: %w", id, err)
		}
		if p.Sold {
			return fmt.Errorf("product %d has already been sold", id)
		}

		p.OrderID = &orderID
		p.Sold = true
		if err := s.productService.UpdateProduct(p); err != nil {
			return fmt.Errorf("product %d orderID cannot be updated: %w", id, err)
		}
	}

	// Publicar evento en el Bus
	eventData := map[string]interface{}{
		"content": fmt.Sprintf("Se ha creado un pedido con el id: %d", orderID),
		"seen":    false,
		"userID":  order.UserID,
	}
	s.eventBusService.Publish("order.created", eventData)

	return nil
}
func (s *OrderService) CheckOrder(order *entities.Order, productsIDs []uint) error {
	if _, err := s.userService.GetUserByID(order.UserID); err != nil {
		log.Printf("Error al verificar usuario %d: %v", order.UserID, err)
		return err
	}

	if _, err := s.addressService.GetAddressByID(order.AddressID); err != nil {
		log.Printf("Error al verificar dirección %d: %v", order.AddressID, err)
		return err
	}

	if _, err := s.carrierService.GetCarrierByID(order.CarrierID); err != nil {
		log.Printf("Error al verificar transportista %d: %v", order.CarrierID, err)
		return err
	}

	for _, productID := range productsIDs {
		product, err := s.productService.GetProductByID(productID)
		if err != nil {
			log.Printf("Error al verificar producto %d: %v", productID, err)
			return err
		}
		if product.Sold {
			return fmt.Errorf("product has already been sold")
		}
	}
	return nil
}

func (s *OrderService) CreateOrder(order *entities.Order, productsIDs []uint) error {
	if err := s.CheckOrder(order, productsIDs); err != nil {
		return err
	}

	orderID, err := s.repo.CreateOrder(order)
	if err != nil {
		return err
	}

	if err := s.AddOrderIDToProducts(orderID, productsIDs); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) UpdateOrder(o *entities.Order) error {
	return s.repo.UpdateOrder(o)
}
func (s *OrderService) DeleteOrder(id uint) error {
	return s.repo.DeleteOrder(id)
}
func (s *OrderService) GetOrderByID(id uint) (*entities.Order, error) {
	return s.repo.GetOrderByID(id)
}

func (s *OrderService) GetFilteredOrders(filters map[string]string) (*[]entities.Order, error) {
	return s.repo.GetFilteredOrders(filters)
}

// Suscribir al evento
func (s *OrderService) ListenPaymentCreated() {
	err := s.eventBusService.Subscribe("payment.created", func(data interface{}) {
		fmt.Println("Evento recibido en OrderService:", data)

		eventData, ok := data.(map[string]interface{})
		if !ok {
			fmt.Println("Error al procesar el evento")
			return
		}

		shippingMap, ok := eventData["shipping"].(map[string]interface{})
		if !ok {
			fmt.Printf("Error: no se puede convertir 'shipping' a map[string]interface{}. Tipo recibido: %T\n", eventData["shipping"])
			return
		}

		// Instanciar directamente la entidad Address usando los datos extraídos del mapa
		// Concatenando calle y número si tu constructor espera 4 campos principales
		streetFull := fmt.Sprintf("%v %v", shippingMap["street"], shippingMap["number"])

		addressEntity := addressEntities.NewAddress(
			streetFull,
			fmt.Sprintf("%v", shippingMap["city"]),
			fmt.Sprintf("%v", shippingMap["province"]),
			fmt.Sprintf("%v", shippingMap["postal_code"]),
		)
		// Crear dirección pasando la Entidad (no el DTO)
		createdAddress, err := s.addressService.CreateAddress(addressEntity)
		if err != nil {
			fmt.Println("Error al crear la dirección:", err)
			return
		}
		// Obtener los IDs directamente como uint
		userID, ok := eventData["userID"].(uint)
		if !ok {
			fmt.Printf("Error: userID no es uint. Tipo recibido: %T\n", eventData["userID"])
			return
		}

		carrierID, ok := eventData["carrierID"].(uint)
		if !ok {
			fmt.Printf("Error: carrierID no es uint. Tipo recibido: %T\n", eventData["carrierID"])
			return
		}

		productID, ok := eventData["productID"].(uint)
		if !ok {
			fmt.Printf("Error: productID no es uint. Tipo recibido: %T\n", eventData["productID"])
			return
		}
		transactionID, ok := eventData["transactionID"].(uint)
		if !ok {
			fmt.Printf("Error: transactionID no es uint. Tipo recibido: %T\n", eventData["transactionID"])
			return
		}

		productIds := []uint{productID}

		orderEntity := entities.NewOrder("pagado", userID, createdAddress.ID, carrierID, transactionID)
		err = s.CreateOrder(orderEntity, productIds)
		if err != nil {
			log.Printf(" Error al crear la orden para el usuario %v: %v\n", userID, err)
			return
		}

		log.Printf(" Orden creada exitosamente \n"+
			" Producto ID: %v\n"+
			" Transportista ID: %v\n"+
			" Dirección ID: %v\n"+
			" Usuario ID: %v\n",
			productID, carrierID, createdAddress.ID, userID,
		)
	})

	if err != nil {
		fmt.Println("Error al suscribirse al evento:", err)
	}
}

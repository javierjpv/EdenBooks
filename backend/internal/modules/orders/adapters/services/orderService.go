package services

import (
	"fmt"
	"log"

	"github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	addressServices "github.com/javierjpv/edenBooks/internal/modules/addresses/domain/services"
	carrierServices "github.com/javierjpv/edenBooks/internal/modules/carriers/domain/services"
	orderDTO "github.com/javierjpv/edenBooks/internal/modules/orders/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/repositories"
	productDTO "github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	productServices "github.com/javierjpv/edenBooks/internal/modules/products/domain/services"
	userService "github.com/javierjpv/edenBooks/internal/modules/users/domain/services"
	eventBusService "github.com/javierjpv/edenBooks/internal/shared/domain/services"
)

type OrderService struct {
	repo            repositories.OrderRepository
	productService  productServices.ProductService
	userService     userService.UserService
	addressService  addressServices.AddressService
	carrierService  carrierServices.CarrierService
	eventBusService eventBusService.EventBus
}

func NewOrderService(repo repositories.OrderRepository, productService productServices.ProductService, addressService addressServices.AddressService, carrierService carrierServices.CarrierService, userService userService.UserService, eventBusService eventBusService.EventBus) *OrderService {
	return &OrderService{repo: repo, productService: productService, addressService: addressService, carrierService: carrierService, userService: userService, eventBusService: eventBusService}
}

func (s *OrderService) AddOrderIDToProducts(orderID uint, productsIDs []uint) error {
	order, err := s.repo.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("order does not exist")
	}

	for _, productID := range productsIDs {
		product, err := s.productService.GetProductByID(productID)
		if err != nil {
			return fmt.Errorf("product does not exist")
		}
		if product.Sold {
			return fmt.Errorf("product has already been sold")
		}
	}
	for _, productsID := range productsIDs {
		product, err := s.productService.GetProductByID(productsID)
		if err != nil {
			return fmt.Errorf("product does not exist")
		}
		product.OrderID = &orderID
		product.Sold = true
		fmt.Println("ProductRequest en AddOrderIDToProducts:", product)
		productRequest := productDTO.NewProductRequest(product.Name, product.Description, product.Price, product.CategoryID, product.UserID, product.ImageURL, product.Sold)
		if err := s.productService.UpdateProduct(product.ID, productRequest); err != nil {
			return fmt.Errorf("product orderID can not be updated")
		}
	}

	userId := order.UserID
	// Publicar evento en el Bus
	eventData := map[string]interface{}{
		"content": fmt.Sprintf("Se ha creado un pedido con el id: %d", orderID),
		"seen":    false,
		"userID":  userId,
	}
	fmt.Println("📢 Publicando evento 'order.created' con datos:", eventData)
	s.eventBusService.Publish("order.created", eventData)

	return nil
}

func (s *OrderService) CheckOrder(o orderDTO.OrderRequest, productsIDs []uint) error {

	if _, err := s.userService.GetUserByID(o.UserID); err != nil {
		log.Printf("Error al verificar el usuario con ID: %d. Error: %v", o.UserID, err)
		return err
	}

	if _, err := s.addressService.GetAddressByID(o.AddressID); err != nil {
		log.Printf("Error al verificar la dirección con ID: %d. Error: %v", o.AddressID, err)
		return err
	}

	if _, err := s.carrierService.GetCarrierByID(o.CarrierID); err != nil {
		log.Printf("Error al verificar el transportista con ID: %d. Error: %v", o.CarrierID, err)
		return err
	}

	//checkTransation???
	for _, productID := range productsIDs {
		product, err := s.productService.GetProductByID(productID)
		if err != nil {
			log.Printf("Error al verificar el producto con ID: %d. Error: %v", productID, err)
			return err
		}
		if product.Sold {
			return fmt.Errorf("product has already been sold")
		}
	}
	return nil
}

func (s *OrderService) CreateOrder(o orderDTO.OrderRequest, productsIDs []uint) error {

	if err := s.CheckOrder(o, productsIDs); err != nil {
		return err
	}

	order := entities.NewOrder(o.State, o.UserID, o.AddressID, o.CarrierID, o.TransactionID)
	orderID, err := s.repo.CreateOrder(order)
	if err != nil {
		return err
	}
	if err := s.AddOrderIDToProducts(orderID, productsIDs); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) UpdateOrder(id uint, o orderDTO.OrderRequest) error {
	order, err := s.repo.GetOrderByID(id)
	if err != nil {
		return err
	}
	order.State = o.State
	order.UserID = o.UserID
	order.AddressID = o.AddressID
	order.CarrierID = o.CarrierID
	return s.repo.UpdateOrder(order)
}
func (s *OrderService) DeleteOrder(id uint) error {
	if _, err := s.repo.GetOrderByID(id); err != nil {
		return err
	}
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

		shippingData, ok := eventData["shipping"].(dto.AddressRequest)
		if !ok {
			fmt.Printf("Error: no se puede convertir 'shipping' a AddressDTO. Tipo recibido: %T\n", eventData["shipping"])
			return
		}

		// Crear el DTO de dirección usando los datos del shipping
		addressDto := dto.NewAddressRequest(
			shippingData.City,
			shippingData.Province,
			shippingData.PostalCode,
			shippingData.Country,
			shippingData.Street,
			shippingData.Number,
		)

		// Crear dirección
		createdAddress, err := s.addressService.CreateAddress(addressDto)
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

		orderDto := orderDTO.NewOrderRequest("pagado", userID, createdAddress.ID, carrierID, transactionID)
		productIds := []uint{productID}

		err = s.CreateOrder(*orderDto, productIds)
		if err != nil {
			log.Printf("❌ Error al crear la orden para el usuario %v: %v\n", userID, err)
			return
		}

		log.Printf("✅ Orden creada exitosamente 🎉\n"+
			"📦 Producto ID: %v\n"+
			"🛵 Transportista ID: %v\n"+
			"📍 Dirección ID: %v\n"+
			"👤 Usuario ID: %v\n",
			productID, carrierID, createdAddress.ID, userID,
		)
	})

	if err != nil {
		fmt.Println("Error al suscribirse al evento:", err)
	}
}

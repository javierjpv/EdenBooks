package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	addresstDTO "github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	orderDTO "github.com/javierjpv/edenBooks/internal/modules/orders/application/dto"
	orderUsecase "github.com/javierjpv/edenBooks/internal/modules/orders/application/useCases"
	productDTO "github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	transactionDTO "github.com/javierjpv/edenBooks/internal/modules/transactions/application/dto"
	transactionUsecase "github.com/javierjpv/edenBooks/internal/modules/transactions/application/useCases"
	eventBusService "github.com/javierjpv/edenBooks/internal/shared/domain/services"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/webhook"
)

type CreateCheckoutSessionRequest struct {
	Product    productDTO.ProductRequest  `json:"product"`
	Shipping   addresstDTO.AddressRequest `json:"shipping"` //coger shipping y llamar al usecase de addres creas el addres solo si no existe y te devuelve el id
	UserID     uint                       `json:"userID"`
	CarrierID  uint                       `json:"carrierID"`
	ProductID  uint                       `json:"productID"`
	SuccessURL string                     `json:"successUrl"`
	CancelURL  string                     `json:"cancelUrl"`
}

type StripeHandler struct {
	eventBusService    eventBusService.EventBus
	orderUsecase       *orderUsecase.OrderUseCase
	transactionUseCase *transactionUsecase.TransactionUseCase
}

func NewStripeHandler(eventBusService eventBusService.EventBus, orderUsecase *orderUsecase.OrderUseCase, transactionUseCase *transactionUsecase.TransactionUseCase) *StripeHandler {
	return &StripeHandler{eventBusService: eventBusService, orderUsecase: orderUsecase, transactionUseCase: transactionUseCase}
}

func (h *StripeHandler) CreateCheckoutSession(c echo.Context) error {
	var req CreateCheckoutSessionRequest
	if err := c.Bind(&req); err != nil {
		log.Println("Error al decodificar JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}

	order := orderDTO.NewOrderRequest("pagado", req.UserID, uint(1), req.CarrierID, uint(1))
	productIds := []uint{req.ProductID}

	if err := h.orderUsecase.CheckOrder(*order, productIds); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	lineItem := &stripe.CheckoutSessionLineItemParams{
		PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency: stripe.String("eur"),
			ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
				Name:        stripe.String(req.Product.Name),
				Description: stripe.String(req.Product.Description),
				Images:      []*string{stripe.String(req.Product.ImageURL)},
			},
			UnitAmount: stripe.Int64(int64(req.Product.Price * 100)),
		},
		Quantity: stripe.Int64(1),
	}
	if req.Product.ImageURL == "" {
		log.Println("Advertencia: req.Product.ImageURL está vacío")
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems:  []*stripe.CheckoutSessionLineItemParams{lineItem},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(req.SuccessURL),
		CancelURL:  stripe.String(req.CancelURL),
	}
	shippingJSON, _ := json.Marshal(req.Shipping)
	params.Params.Metadata = map[string]string{
		"total":     fmt.Sprintf("%.2f", req.Product.Price),
		"shipping":  string(shippingJSON), // Esto puede necesitar ser ajustado
		"userID":    strconv.Itoa(int(req.UserID)),
		"carrierID": strconv.Itoa(int(req.CarrierID)),
		"productID": strconv.Itoa(int(req.ProductID)),
	}

	session, err := session.New(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"sessionId": session.ID})
}

type HandleStripeWebhookRequest struct {
	Total     float64                    `json:"total"`
	Shipping  addresstDTO.AddressRequest `json:"shipping"`
	UserID    uint                       `json:"userID"`
	CarrierID uint                       `json:"carrierID"`
	ProductID uint                       `json:"productID"`
}

func (h *StripeHandler) HandleStripeWebhook(c echo.Context) error {
	// Leer el cuerpo sin procesarlo
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("❌ Error leyendo el body: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Cargar el secreto del webhook desde el .env
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if endpointSecret == "" {
		log.Println("⚠️ STRIPE_WEBHOOK_SECRET no está configurado en el .env")
		return c.JSON(http.StatusInternalServerError, "Webhook secret not configured")
	}

	// Obtener la firma del header
	stripeSignature := c.Request().Header.Get("Stripe-Signature")

	// Verificar la firma y construir el evento
	event, err := webhook.ConstructEvent(body, stripeSignature, endpointSecret)
	if err != nil {
		log.Printf("❌ Error verificando la firma del webhook: %v", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Manejar el evento
	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			log.Printf("❌ Error parseando el evento: %v", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if session.Metadata == nil {
			log.Println("❌ Metadata no presente en el webhook")
			return c.JSON(http.StatusBadRequest, "Missing metadata")
		}

		// Verificar la conversión de los campos
		// Mostrar metadata antes de la conversión
		log.Printf("Metadata recibido: %v", session.Metadata)

		// Verificar la conversión de los campos
		total, err := strconv.ParseFloat(session.Metadata["total"], 64)
		if err != nil {
			log.Printf("❌ Error parseando 'total': %v", err)
			return c.JSON(http.StatusBadRequest, "Invalid total data")
		}

		// Convertir userID, carrierID y productID a uint
		userID, err := strconv.ParseUint(session.Metadata["userID"], 10, 32)
		if err != nil {
			log.Printf("❌ Error parseando 'userID': %v. Valor recibido: %v", err, session.Metadata["userID"])
			return c.JSON(http.StatusBadRequest, "Invalid userID")
		}
		log.Printf("userID convertido: %d", userID)

		carrierID, err := strconv.ParseUint(session.Metadata["carrierID"], 10, 32)
		if err != nil {
			log.Printf("❌ Error parseando 'carrierID': %v. Valor recibido: %v", err, session.Metadata["carrierID"])
			return c.JSON(http.StatusBadRequest, "Invalid carrierID")
		}
		log.Printf("carrierID convertido: %d", carrierID)

		productID, err := strconv.ParseUint(session.Metadata["productID"], 10, 32)
		if err != nil {
			log.Printf("❌ Error parseando 'productID': %v. Valor recibido: %v", err, session.Metadata["productID"])
			return c.JSON(http.StatusBadRequest, "Invalid productID")
		}
		log.Printf("productID convertido: %d", productID)

		// Convertir shipping de string a objeto JSON
		var shipping addresstDTO.AddressRequest
		if err := json.Unmarshal([]byte(session.Metadata["shipping"]), &shipping); err != nil {
			log.Printf("❌ Error parseando shipping: %v", err)
			return c.JSON(http.StatusBadRequest, "Invalid shipping data")
		}

		log.Printf("Shipping convertido: %+v", shipping)

		//Publicar evento en el Bus
		transaction := transactionDTO.NewTransactionDTO("tarjeta", float64(total))
		createdTransaction, err := h.transactionUseCase.CreateTransaction(*transaction)
		if err != nil {
			return err
		}
		createdTransactionID := createdTransaction.ID

		eventData := map[string]interface{}{
			"shipping":      shipping,
			"userID":        uint(userID),
			"carrierID":     uint(carrierID),
			"productID":     uint(productID),
			"transactionID": uint(createdTransactionID),
		}
		h.eventBusService.Publish("payment.created", eventData)

	case "checkout.session.expired":
		log.Println("⚠️ La sesión de pago ha expirado")
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "received"})
}

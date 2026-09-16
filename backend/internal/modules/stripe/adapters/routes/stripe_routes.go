package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/stripe/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterStripeRoutes(e *echo.Echo, stripeHandler *handlers.StripeHandler)  {
	stripeGroup:=e.Group("/stripe")
	stripeGroup.POST("",stripeHandler.CreateCheckoutSession)
	stripeGroup.POST("/webhook",stripeHandler.HandleStripeWebhook)

}
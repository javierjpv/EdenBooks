package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/addresses/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterAddressRoutes(e *echo.Echo, addressHandler *handlers.AddressHandler)  {
	addressGroup:=e.Group("/addresses")
	addressGroup.POST("",addressHandler.CreateAddress)
	addressGroup.PUT("/:id",addressHandler.UpdateAddress)
	addressGroup.GET("",addressHandler.GetFilteredAddresses)
	addressGroup.GET("/:id",addressHandler.GetAddress)
}
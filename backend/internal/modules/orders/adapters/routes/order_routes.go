package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/orders/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterOrderRoutes(e *echo.Echo, orderHandler *handlers.OrderHandler)  {
	orderGroup:=e.Group("/orders")
	// orderGroup.POST("",orderHandler.CreateOrder)
    // orderGroup.PUT("/:id",orderHandler.UpdateOrder)
	// orderGroup.DELETE("/:id",orderHandler.DeleteOrder)
	// orderGroup.GET("/:id",orderHandler.GetOrderByID)
	orderGroup.GET("", orderHandler.GetFilteredOrders)
}

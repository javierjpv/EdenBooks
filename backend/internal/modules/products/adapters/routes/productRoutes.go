package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/products/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterProductRoutes(e *echo.Echo, productHandler *handlers.ProductHandler)  {
	productGroup:=e.Group("/products")
	productGroup.POST("",productHandler.CreateProduct)
    productGroup.PUT("/:id",productHandler.UpdateProduct)
	productGroup.DELETE("/:id",productHandler.DeleteProduct)
	productGroup.GET("/:id",productHandler.GetProductByID)
	productGroup.GET("", productHandler.GetProductsWithFavorites)
	productGroup.POST("/:id/favorite", productHandler.AddToFavorites)
    productGroup.DELETE("/:id/favorite", productHandler.RemoveFromFavorites)
	productGroup.GET("/favorites", productHandler.GetFavorites)
	productGroup.GET("/withfavorites", productHandler.GetProductsWithFavorites)
}


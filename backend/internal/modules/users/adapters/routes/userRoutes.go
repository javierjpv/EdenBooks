package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/users/adapters/handlers"
	"github.com/labstack/echo/v4"
)


func RegisterUserRoutes(e *echo.Echo, userHandler *handlers.UserHandler){
	userGroup:=e.Group("/users")
	userGroup.POST("/register",userHandler.Register)
	userGroup.POST("/login",userHandler.Login)
	userGroup.GET("/:id",userHandler.GetUserByID)
	userGroup.PUT("/:id",userHandler.UpdateUser)
}
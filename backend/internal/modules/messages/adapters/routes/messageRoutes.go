package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/messages/adapters/handlers"
	"github.com/labstack/echo/v4"
)


func RegisterMessageRoutes(e *echo.Echo, messageHandler *handlers.MessageHandler)  {
	messageGroup:=e.Group("/messages")
	messageGroup.POST("",messageHandler.CreateMessage)
    messageGroup.PUT("/:id",messageHandler.UpdateMessage)
	messageGroup.DELETE("/:id",messageHandler.DeleteMessage)
	messageGroup.GET("/:id",messageHandler.GetMessageByID)
	messageGroup.GET("",messageHandler.GetFilteredMessages)
}
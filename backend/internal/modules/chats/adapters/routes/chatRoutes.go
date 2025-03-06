package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/chats/adapters/handlers"
	"github.com/labstack/echo/v4"
)


func RegisterChatRoutes(e *echo.Echo, chatHandler *handlers.ChatHandler)  {
	chatGroup:=e.Group("/chats")
	chatGroup.POST("",chatHandler.CreateChatWithUsers)
    chatGroup.PUT("/:id",chatHandler.UpdateChat)
	chatGroup.DELETE("/:id",chatHandler.DeleteChat)
	chatGroup.GET("/:id",chatHandler.GetChatByID)
	chatGroup.GET("",chatHandler.GetFilteredChats)
}
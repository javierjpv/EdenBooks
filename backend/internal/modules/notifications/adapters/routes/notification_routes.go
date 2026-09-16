package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/notifications/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterNotificationRoutes(e *echo.Echo, notificationHandler *handlers.NotificationHandler)  {
	notificationGroup:=e.Group("/notifications")
	// notificationGroup.POST("",notificationHandler.CreateNotification)
    // notificationGroup.PUT("/:id",notificationHandler.UpdateNotification)
	// notificationGroup.DELETE("/:id",notificationHandler.DeleteNotification)
	// notificationGroup.GET("/:id",notificationHandler.GetNotificationByID)
	notificationGroup.GET("", notificationHandler.GetFilteredNotifications)
	
}
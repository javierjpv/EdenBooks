package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/javierjpv/edenBooks/internal/modules/notifications/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/notifications/application/useCases"
	"github.com/javierjpv/edenBooks/internal/shared/auth"
	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	useCase usecases.NotificationUseCase
}

func NewNotificationHandler(useCase usecases.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{useCase: useCase}
}

func (h *NotificationHandler) CreateNotification(c echo.Context) error {
	var notification dto.NotificationRequest

	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	if err := h.useCase.CreateNotification(notification); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear el notification"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "notification creado correctamente"})
}

func (h *NotificationHandler) UpdateNotification(c echo.Context) error {
	var notification dto.NotificationRequest
	notificationID := c.Param("id")
	id, err := strconv.Atoi(notificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	if err := h.useCase.UpdateNotification(uint(id), notification); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el notification"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "notification actualizado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}

func (h *NotificationHandler) DeleteNotification(c echo.Context) error {
	notificationID := c.Param("id")
	id, err := strconv.Atoi(notificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := h.useCase.DeleteNotification(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar el notification"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "notification borrado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}

func (h *NotificationHandler) GetNotificationByID(c echo.Context) error {
	notificationID := c.Param("id")
	id, err := strconv.Atoi(notificationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	notification, err := h.useCase.GetNotificationByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener el notification"})
	}
	return c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandler) GetFilteredNotifications(c echo.Context) error {
	fmt.Println("Procesando solicitud...")
	// Obtener el token del encabezado Authorization
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := auth.ExtractUserIDFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido: " + err.Error()})
	}

	fmt.Println("user_id", userID)
	// Extraer filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}
	// Agregar el userID como un filtro obligatorio
	filters["user_id"] = fmt.Sprintf("%d", userID)

	notifications, err := h.useCase.GetFilteredNotifications(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, notifications)
}

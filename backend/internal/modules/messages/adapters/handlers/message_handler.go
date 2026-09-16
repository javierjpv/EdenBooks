package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/messages/application/useCases"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	useCase usecases.MessageUseCase
}

func NewMessageHandler(useCase usecases.MessageUseCase) *MessageHandler {
	return &MessageHandler{useCase: useCase}
}

func (h *MessageHandler) CreateMessage(c echo.Context) error {
	var message dto.MessageRequest

	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	if _, err := h.useCase.CreateMessage(message); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la transaccion"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "transaccion creada correctamente"})
}

func (h *MessageHandler) UpdateMessage(c echo.Context) error {
	var message dto.MessageRequest
	messageID := c.Param("id")
	id, err := strconv.Atoi(messageID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	if err := h.useCase.UpdateMessage(uint(id), message); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la transaccion"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "transaccion actualizada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}

func (h *MessageHandler) DeleteMessage(c echo.Context) error {
	messageID := c.Param("id")
	id, err := strconv.Atoi(messageID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := h.useCase.DeleteMessage(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar la transaccion"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "transaccion borrada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}

func (h *MessageHandler) GetMessageByID(c echo.Context) error {
	messageID := c.Param("id")
	id, err := strconv.Atoi(messageID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	message, err := h.useCase.GetMessageByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener la transaccion"})
	}
	return c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) GetFilteredMessages(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	messages, err := h.useCase.GetFilteredMessages(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, messages)
}

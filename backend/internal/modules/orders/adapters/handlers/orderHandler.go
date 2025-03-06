package handlers

import (

	"fmt"
	"net/http"
	"strconv"
	"strings"


	"github.com/javierjpv/edenBooks/internal/modules/orders/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/orders/application/useCases"
	"github.com/javierjpv/edenBooks/internal/shared/auth"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	useCase usecases.OrderUseCase
}

func NewOrderHandler(useCase usecases.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: useCase}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req struct {
		OrderDTO   dto.OrderDTO `json:"order"`
		ProductIDs []uint       `json:"product_ids"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	if len(req.ProductIDs) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Debe incluir al menos un producto en la orden"})
	}
	if err := h.useCase.CreateOrder(req.OrderDTO, req.ProductIDs); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear el order"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "order creado correctamente"})
}

func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	var order dto.OrderDTO
	orderID := c.Param("id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}
	if err := h.useCase.UpdateOrder(uint(id), order); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el order"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "order actualizado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}

func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	orderID := c.Param("id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	if err := h.useCase.DeleteOrder(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar el order"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "order borrado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}

func (h *OrderHandler) GetOrderByID(c echo.Context) error {
	orderID := c.Param("id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	order, err := h.useCase.GetOrderByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener el order"})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetFilteredOrders(c echo.Context) error {
	fmt.Println("Procesando solicitud...");
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

	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}
	// Agregar el userID como un filtro obligatorio
	filters["user_id"] = fmt.Sprintf("%d", userID)
	orders, err := h.useCase.GetFilteredOrders(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, orders)
}


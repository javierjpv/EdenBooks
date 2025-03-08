package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/carriers/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/carriers/application/useCases"
	"github.com/labstack/echo/v4"
)

type CarrierHandler struct {
	useCase usecases.CarrierUseCase
}

func NewCarrierHandler(useCase usecases.CarrierUseCase) *CarrierHandler {
	return &CarrierHandler{useCase: useCase}
}

func (h *CarrierHandler) CreateCarrier(c echo.Context) error {
	var carrier dto.CarrierRequest
	//se pasa como un puntero ya que esto hace que pase de ser un json a ser un objeto de go
	if err := c.Bind(&carrier); err != nil {
		return c.String(http.StatusBadRequest, "error al recoger los datos")
	}

	if err := h.useCase.CreateCarrier(carrier); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la transportista"})
	}
	return c.String(http.StatusCreated, "transportista creado correctamente")
}
func (h *CarrierHandler) UpdateCarrier(c echo.Context) error {
	var carrier dto.CarrierRequest
	carrierID := c.Param("id")
	id, err := strconv.Atoi(carrierID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	//se pasa como un puntero ya que esto hace que pase de ser un json a ser un objeto de go
	if err := c.Bind(&carrier); err != nil {
		return c.String(http.StatusBadRequest, "error al recoger los datos")
	}

	if err := h.useCase.UpdateCarrier(uint(id), carrier); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la transportista"})
	}
	return c.String(http.StatusCreated, "transportista actualizado correctamente")
}
func (h *CarrierHandler) GetCarrier(c echo.Context) error {
	carrierID := c.Param("id")
	id, err := strconv.Atoi(carrierID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	carrier, err := h.useCase.GetCarrierByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transportista no encontrado"})
	}
	return c.JSON(http.StatusOK, carrier)
}

func (h *CarrierHandler) GetFilteredCarrieres(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	carrieres, err := h.useCase.GetFilteredCarrieres(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, carrieres)
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/addresses/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/addresses/application/useCases"
	"github.com/labstack/echo/v4"
)

type AddressHandler struct {
	useCase usecases.AddressUseCase
}

func NewAddressHandler(useCase usecases.AddressUseCase) *AddressHandler {
	return &AddressHandler{useCase: useCase}
}

func (h *AddressHandler) CreateAddress(c echo.Context) error {
	var address dto.AddressRequest
	//se pasa como un puntero ya que esto hace que pase de ser un json a ser un objeto de go
	if err := c.Bind(&address); err != nil {
		return c.String(http.StatusBadRequest, "error al recoger los datos")
	}
	createdAddress, err := h.useCase.CreateAddress(address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la dirección"})
	}
	return c.JSON(http.StatusCreated, createdAddress)
}
func (h *AddressHandler) UpdateAddress(c echo.Context) error {
	var address dto.AddressRequest
	addressID := c.Param("id")
	id, err := strconv.Atoi(addressID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}
	//se pasa como un puntero ya que esto hace que pase de ser un json a ser un objeto de go
	if err := c.Bind(&address); err != nil {
		return c.String(http.StatusBadRequest, "error al recoger los datos")
	}

	if err := h.useCase.UpdateAddress(uint(id), address); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la dirección"})
	}
	return c.String(http.StatusCreated, "Direccion actualizada correctamente")
}
func (h *AddressHandler) GetAddress(c echo.Context) error {
	addressID := c.Param("id")
	id, err := strconv.Atoi(addressID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	address, err := h.useCase.GetAddressByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Dirección no encontrada"})
	}
	return c.JSON(http.StatusOK, address)
}

func (h *AddressHandler) GetFilteredAddresses(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	addresses, err := h.useCase.GetFilteredAddresses(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, addresses)
}

// curl -X GET "http://localhost:6969/addresses?city=Madrid&sort_by=created_at&order=asc&page=1&limit=2"      -H "Content-Type: application/json" | jq

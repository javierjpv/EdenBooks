package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/categories/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/categories/application/useCases"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	useCase usecases.CategoryUseCase
}

func NewCategoryHandler(useCase usecases.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{useCase: useCase}
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var category dto.CategoryRequest
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error al recoger los datos"})
	}

	if err := h.useCase.CreateCategory(category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la categoria"})
	}
	return c.String(http.StatusCreated, "Categoria creada correctamente")

}

func (h *CategoryHandler) GetCategory(c echo.Context) error {
	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
	}

	category, err := h.useCase.GetCategoryByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Categoria no encontrada"})
	}
	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetAllCategories(c echo.Context) error {
	categories, err := h.useCase.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Categorias no encontradas"})
	}
	return c.JSON(http.StatusOK, categories)
}

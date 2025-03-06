package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/products/application/useCases"
	"github.com/labstack/echo/v4"
)


type ProductHandler struct{
	useCase usecases.ProductUseCase
}

func NewProductHandler(useCase usecases.ProductUseCase)*ProductHandler{
	return &ProductHandler{useCase: useCase}
}

func (h *ProductHandler)CreateProduct(c echo.Context)error{
	var product dto.ProductDTO
	
	if err:=c.Bind(&product);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if err:=h.useCase.CreateProduct(product);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear el producto"})
	}
	return c.JSON(http.StatusCreated,map[string]string{"message": "producto creado correctamente"})
}

func (h *ProductHandler)UpdateProduct(c echo.Context)error{
	var product dto.ProductDTO
	productID:=c.Param("id")
    id, err := strconv.Atoi(productID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	if err:=c.Bind(&product);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if err:=h.useCase.UpdateProduct(uint(id),product);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar el producto"})
	}
	return c.JSON(http.StatusOK,map[string]string{"message": "producto actualizado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}


func (h *ProductHandler)DeleteProduct(c echo.Context)error{
	productID:=c.Param("id")
    id, err := strconv.Atoi(productID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	
	if err:=h.useCase.DeleteProduct(uint(id));err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar el producto"})
	}
	return c.JSON(http.StatusOK,map[string]string{"message": "producto borrado correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}


func (h *ProductHandler)GetProductByID(c echo.Context)error{
	productID:=c.Param("id")
    id, err := strconv.Atoi(productID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	product,err:=h.useCase.GetProductByID(uint(id));
	if err!=nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener el producto"})
	}
	return c.JSON(http.StatusOK, product)
}
func (h *ProductHandler) GetFilteredProducts(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	products, err := h.useCase.GetFilteredProducts(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, products)
}

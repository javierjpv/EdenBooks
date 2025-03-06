package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/reviews/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/reviews/application/useCases"
	"github.com/labstack/echo/v4"
)


type ReviewHandler struct{
	useCase usecases.ReviewUseCase
}

func NewReviewHandler(useCase usecases.ReviewUseCase)*ReviewHandler{
	return &ReviewHandler{useCase: useCase}
}

func (h *ReviewHandler)CreateReview(c echo.Context)error{
	var review dto.ReviewDTO
	
	if err:=c.Bind(&review);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if err:=h.useCase.CreateReview(review);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la reseña"})
	}
	return c.JSON(http.StatusCreated,map[string]string{"message": "reseña creada correctamente"})
}

func (h *ReviewHandler)UpdateReview(c echo.Context)error{
	var review dto.ReviewDTO
	reviewID:=c.Param("id")
    id, err := strconv.Atoi(reviewID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	if err:=c.Bind(&review);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if err:=h.useCase.UpdateReview(uint(id),review);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la reseña"})
	}
	return c.JSON(http.StatusOK,map[string]string{"message": "reseña actualizada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}


func (h *ReviewHandler)DeleteReview(c echo.Context)error{
	reviewID:=c.Param("id")
    id, err := strconv.Atoi(reviewID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	
	if err:=h.useCase.DeleteReview(uint(id));err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar la reseña"})
	}
	return c.JSON(http.StatusOK,map[string]string{"message": "reseña borrada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}


func (h *ReviewHandler)GetReviewByID(c echo.Context)error{
	reviewID:=c.Param("id")
    id, err := strconv.Atoi(reviewID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	review,err:=h.useCase.GetReviewByID(uint(id));
	if err!=nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener la reseña"})
	}
	return c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) GetFilteredReviews(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	reviews, err := h.useCase.GetFilteredReviews(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, reviews)
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/transactions/application/dto"
	usecases "github.com/javierjpv/edenBooks/internal/modules/transactions/application/useCases"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct{
	useCase usecases.TransactionUseCase
}

func NewTransactionHandler(useCase usecases.TransactionUseCase)*TransactionHandler{
	return &TransactionHandler{useCase: useCase}
}

func (h *TransactionHandler)CreateTransaction(c echo.Context)error{
	var transaction dto.TransactionDTO
	
	if err:=c.Bind(&transaction);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if _,err:=h.useCase.CreateTransaction(transaction);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al crear la transaccion"})
	}
	return c.JSON(http.StatusCreated,map[string]string{"message": "transaccion creada correctamente"})
}

func (h *TransactionHandler)UpdateTransaction(c echo.Context)error{
	var transaction dto.TransactionDTO
	transactionID:=c.Param("id")
    id, err := strconv.Atoi(transactionID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	if err:=c.Bind(&transaction);err!=nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error al recoger los datos"})
	}

	
	if err:=h.useCase.UpdateTransaction(uint(id),transaction);err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al actualizar la transaccion"})
	}
	return c.JSON(http.StatusOK,map[string]string{"message": "transaccion actualizada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien

}


func (h *TransactionHandler)DeleteTransaction(c echo.Context)error{
	transactionID:=c.Param("id")
    id, err := strconv.Atoi(transactionID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	
	if err:=h.useCase.DeleteTransaction(uint(id));err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error al borrar la transaccion"})
	}
	return c.JSON(http.StatusOK,map[string]string{"message": "transaccion borrada correctamente"})
	// return c.NoContent(http.StatusNoContent) convencion rest, aunque si quieres enviar un mensaje esta bien
}


func (h *TransactionHandler)GetTransactionByID(c echo.Context)error{
	transactionID:=c.Param("id")
    id, err := strconv.Atoi(transactionID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID inválido"})
    }
	transaction,err:=h.useCase.GetTransactionByID(uint(id));
	if err!=nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error al obtener la transaccion"})
	}
	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetFilteredTransactions(c echo.Context) error {
	// Extraer todos los filtros dinámicos de la URL
	filters := make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	transactions, err := h.useCase.GetFilteredTransactions(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}

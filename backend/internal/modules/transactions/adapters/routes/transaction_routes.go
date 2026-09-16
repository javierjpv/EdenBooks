package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/transactions/adapters/handlers"
	"github.com/labstack/echo/v4"
)


func RegisterTransactionRoutes(e *echo.Echo, transactionHandler *handlers.TransactionHandler)  {
	transactionGroup:=e.Group("/transactions")
	transactionGroup.POST("",transactionHandler.CreateTransaction)
    transactionGroup.PUT("/:id",transactionHandler.UpdateTransaction)
	transactionGroup.DELETE("/:id",transactionHandler.DeleteTransaction)
	transactionGroup.GET("/:id",transactionHandler.GetTransactionByID)
	transactionGroup.GET("", transactionHandler.GetFilteredTransactions)
}
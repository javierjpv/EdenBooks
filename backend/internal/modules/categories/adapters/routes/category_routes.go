package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/categories/adapters/handlers"
	"github.com/labstack/echo/v4"
)


func RegisterCategoryRoutes(e *echo.Echo, categoryHandler *handlers.CategoryHandler)  {
	categoryGroup:=e.Group("/categories")
	categoryGroup.POST("",categoryHandler.CreateCategory)
	categoryGroup.GET("",categoryHandler.GetAllCategories)
	categoryGroup.GET("/:id",categoryHandler.GetCategory)
}

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Ficción", "description": "Novelas y relatos imaginarios"}'

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Ciencia", "description": "Libros sobre física, química, biología, etc."}'

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Historia", "description": "Libros sobre eventos históricos y personajes"}'

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Filosofía", "description": "Reflexiones sobre la existencia, el conocimiento y la ética"}'

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Terror", "description": "Relatos de horror y suspenso"}'

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Tecnología", "description": "Libros sobre informática, IA y avances tecnológicos"}'

// curl -X POST http://localhost:6969/categories -H "Content-Type: application/json" -d '{"name": "Autoayuda", "description": "Libros para el crecimiento personal y el bienestar"}'

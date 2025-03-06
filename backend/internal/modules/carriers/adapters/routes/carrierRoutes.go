package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/carriers/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterCarrierRoutes(e *echo.Echo, carrierHandler *handlers.CarrierHandler)  {
	carrierGroup:=e.Group("/carriers")
	carrierGroup.POST("",carrierHandler.CreateCarrier)
	carrierGroup.PUT("/:id",carrierHandler.UpdateCarrier)
	carrierGroup.GET("",carrierHandler.GetFilteredCarrieres)
	carrierGroup.GET("/:id",carrierHandler.GetCarrier)
}

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 1", "contact": "contact1@example.com"}'

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 2", "contact": "contact2@example.com"}'

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 3", "contact": "contact3@example.com"}'

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 4", "contact": "contact4@example.com"}'

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 5", "contact": "contact5@example.com"}'

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 6", "contact": "contact6@example.com"}'

// curl -X POST http://localhost:6969/carriers \
//      -H "Content-Type: application/json" \
//      -d '{"name": "Carrier 7", "contact": "contact7@example.com"}'

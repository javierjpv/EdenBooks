package routes

import (
	"github.com/javierjpv/edenBooks/internal/modules/reviews/adapters/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterReviewRoutes(e *echo.Echo, reviewHandler *handlers.ReviewHandler)  {
	reviewGroup:=e.Group("/reviews")
	reviewGroup.POST("",reviewHandler.CreateReview)
    reviewGroup.PUT("/:id",reviewHandler.UpdateReview)
	reviewGroup.DELETE("/:id",reviewHandler.DeleteReview)
	reviewGroup.GET("/:id",reviewHandler.GetReviewByID)
	reviewGroup.GET("", reviewHandler.GetFilteredReviews) 
}
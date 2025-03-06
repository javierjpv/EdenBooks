package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/reviews/application/dto"
)



type ReviewService interface{
	CreateReview(review dto.ReviewDTO)error

	UpdateReview(id uint, r dto.ReviewDTO)error

	DeleteReview(id uint)error

	GetReviewByID(id uint)(*entities.Review,error)
	
	GetFilteredReviews(filters map[string]string) ([]entities.Review, error) 
}
package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/reviews/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
)

type ReviewService interface {
	CreateReview(review dto.ReviewRequest) error

	UpdateReview(id uint, r dto.ReviewRequest) error

	DeleteReview(id uint) error

	GetReviewByID(id uint) (*entities.Review, error)

	GetFilteredReviews(filters map[string]string) ([]entities.Review, error)
}

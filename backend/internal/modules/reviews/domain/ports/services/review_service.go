package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
)

type ReviewService interface {
	CreateReview(review *entities.Review) error

	UpdateReview(r *entities.Review) error

	DeleteReview(id uint) error

	GetReviewByID(id uint) (*entities.Review, error)

	GetFilteredReviews(filters map[string]string) ([]entities.Review, error)
}

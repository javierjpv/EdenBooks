package repositories

import "github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"

type ReviewRepository interface{
	CreateReview(review *entities.Review)error

	UpdateReview(review *entities.Review)error

	DeleteReview(id uint)error

	GetReviewByID(id uint)(*entities.Review,error)
	
	GetFilteredReviews(filters map[string]string) ([]entities.Review, error)
}
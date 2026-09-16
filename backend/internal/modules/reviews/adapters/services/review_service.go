package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/ports/repositories"
)

type ReviewService struct {
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(r *entities.Review) error {
	return s.repo.CreateReview(r)
}

func (s *ReviewService) UpdateReview(r *entities.Review) error {
	return s.repo.UpdateReview(r)
}
func (s *ReviewService) DeleteReview(id uint) error {
	if _, err := s.repo.GetReviewByID(id); err != nil {
	}
	return s.repo.DeleteReview(id)
}
func (s *ReviewService) GetReviewByID(id uint) (*entities.Review, error) {
	return s.repo.GetReviewByID(id)
}
func (s *ReviewService) GetFilteredReviews(filters map[string]string) ([]entities.Review, error) {
	return s.repo.GetFilteredReviews(filters)
}

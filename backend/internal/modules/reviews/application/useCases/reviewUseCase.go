package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/reviews/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/services"
)

var (
	ErrMissingFields = fmt.Errorf("all fields are required")
	ErrInvalid       = fmt.Errorf("invalid ID")
)

type ReviewUseCase struct {
	service services.ReviewService
}

func NewReviewUseCase(service services.ReviewService) *ReviewUseCase {
	return &ReviewUseCase{service: service}
}

func (u *ReviewUseCase) CreateReview(r dto.ReviewRequest) error {
	if r.Comment == "" {
		return ErrMissingFields
	}
	if r.Rating < 0 {
		return fmt.Errorf("rating can not be < 0")
	}
	if r.ProductID == 0 || r.UserID == 0 {
		return ErrInvalid
	}
	return u.service.CreateReview(r)
}

func (u *ReviewUseCase) UpdateReview(id uint, r dto.ReviewRequest) error {
	if r.Comment == "" {
		return ErrMissingFields
	}
	if r.Rating < 0 {
		return fmt.Errorf("rating can not be < 0")
	}
	if id == 0 || r.ProductID == 0 || r.UserID == 0 {
		return ErrInvalid
	}
	return u.service.UpdateReview(id, r)
}

func (u *ReviewUseCase) DeleteReview(id uint) error {
	if id == 0 {
		return ErrInvalid
	}
	return u.service.DeleteReview(id)
}

func (u *ReviewUseCase) GetReviewByID(id uint) (*dto.ReviewResponse, error) {
	if id == 0 {
		return nil, ErrInvalid
	}
	review, err := u.service.GetReviewByID(id)
	if err != nil {
		return nil, err
	}
	reviewResponse := dto.NewReviewResponse(review.ID,review.CreatedAt,review.UpdatedAt,review.Rating,review.Comment,review.UserID,review.ProductID)
	return reviewResponse, nil
}
func (u *ReviewUseCase) GetFilteredReviews(filters map[string]string) ([]dto.ReviewResponse, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"rating": true, "user_id": true, "product_id": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	// return u.service.GetFilteredReviewes(filters)
	reviewes, err := u.service.GetFilteredReviews(filters)
	if err != nil {
		return nil, err
	}

	// Convertir cada Review a ReviewResponse
	var reviewResponses []dto.ReviewResponse
	for _, review := range reviewes {
		reviewResponses = append(reviewResponses, *dto.NewReviewResponse(
			review.ID,review.CreatedAt,review.UpdatedAt,review.Rating,review.Comment,review.UserID,review.ProductID))
	}

	return reviewResponses, nil
}
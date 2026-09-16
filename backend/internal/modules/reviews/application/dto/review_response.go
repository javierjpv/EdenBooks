package dto

import "time"

type ReviewResponse struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Rating    int
	Comment   string
	UserID    uint
	ProductID uint
}

func NewReviewResponse(id uint, createdAt, updatedAt time.Time, rating int, comment string, userID, productID uint) *ReviewResponse {
	return &ReviewResponse{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Rating:    rating,
		Comment:   comment,
		UserID:    userID,
		ProductID: productID,
	}
}

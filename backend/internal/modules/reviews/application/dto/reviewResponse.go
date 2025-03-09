package dto

import "time"

type ReviewResponse struct {
	//puede colocar la etiqueta json aqui para devolver en el front como quieras cada propiedad
	//por ejemplo en vez de reponder con ID puedes responder con id
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

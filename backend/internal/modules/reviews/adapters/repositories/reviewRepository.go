package repositories

import (
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	"gorm.io/gorm"
)

type ReviewRepository struct{
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB)*ReviewRepository{
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository)CreateReview(review *entities.Review)error{
	if err:=r.db.Create(review).Error; err!=nil {
		return err
	}
	return nil

}
func (r *ReviewRepository)UpdateReview(review *entities.Review) error{
	err:=r.db.Save(review).Error
	if  err!=nil {
		return err
	}
	return nil
}

func (r *ReviewRepository)DeleteReview(id uint) error{
	err:=r.db.Delete(&entities.Review{},id).Error
	if err!=nil {
		return err
	}
	return nil
}

func (r *ReviewRepository)GetReviewByID(id uint) (*entities.Review,error){
	var review entities.Review 
	err:=r.db.First(&review,id).Error
	if err!=nil {
		return nil,err
	}
	return &review,nil
}
func (r *ReviewRepository) GetFilteredReviews(filters map[string]string) ([]entities.Review, error) {
	var reviews []entities.Review
	query := r.db

	// Aplicar filtros dinámicos
	for key, value := range filters {
		switch key {
		case "rating":
			query = query.Where("rating = ?", value)
		case "min_rating":
			query = query.Where("rating >= ?", value)
		case "max_rating":
			query = query.Where("rating <= ?", value)
		case "user_id":
			query = query.Where("user_id = ?", value)
		case "product_id":
			query = query.Where("product_id = ?", value)
		case "comment":
			query = query.Where("comment ILIKE ?", "%"+value+"%") // Búsqueda parcial
		}
	}

	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := filters["order"]
		query = query.Order(sortBy + " " + order)
	}

	// Aplicar paginación si está presente
	limitInt := 50 // Límite por defecto
	if limit, exists := filters["limit"]; exists {
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = parsedLimit
		}
	}
	query = query.Limit(limitInt)

	if page, exists := filters["page"]; exists {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			query = query.Offset((pageInt - 1) * limitInt)
		}
	}

	// Ejecutar la consulta
	if err := query.Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

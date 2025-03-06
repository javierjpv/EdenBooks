package services

import (
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/reviews/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/reviews/domain/repositories"
)


type ReviewService struct{
	repo repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository)*ReviewService{
	return &ReviewService{repo: repo}
}

func(s * ReviewService)CreateReview(r dto.ReviewDTO)error{
	review:=entities.NewReview(r.Rating,r.Comment,r.UserID,r.ProductID)//se creara un user sin order id ya q esta se llevara a cabo mas adelante
    return s.repo.CreateReview(review)
}

func(s * ReviewService)UpdateReview(id uint, r dto.ReviewDTO)error{
	review,err:=s.repo.GetReviewByID(id)
    if err!=nil{
		return err
	}
    review.Rating=r.Rating
	review.Comment=r.Comment
	review.UserID=r.UserID
	review.ProductID=r.ProductID
	return s.repo.UpdateReview(review)	
}
func(s * ReviewService)DeleteReview(id uint)error{
	if _,err:=s.repo.GetReviewByID(id);err!=nil{}
	return s.repo.DeleteReview(id)
}
func(s * ReviewService)GetReviewByID(id uint)(*entities.Review,error){
	return s.repo.GetReviewByID(id)
}
func (s *ReviewService) GetFilteredReviews(filters map[string]string) ([]entities.Review, error) {
	return s.repo.GetFilteredReviews(filters)
}

package entities

import "gorm.io/gorm"


type Review struct{
	gorm.Model
	Rating int
	Comment string
	UserID uint //Relacion 1:N
	ProductID uint //Relacion 1:N

}

func NewReview(rating int, comment string,userID uint,productID uint)*Review{
return &Review{Rating: rating,Comment: comment,UserID: userID,ProductID: productID}
}
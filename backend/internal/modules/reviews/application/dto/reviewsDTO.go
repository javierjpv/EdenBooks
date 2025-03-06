package dto

type ReviewDTO struct{
	Rating int       `json:"rating"`
	Comment string   `json:"comment"`
	UserID uint      `json:"userID"`
	ProductID uint   `json:"productID"`
}
package entities

type ProductWithFavoriteStatus struct {
    Product
    IsFavorite bool `json:"is_favorite"`
}
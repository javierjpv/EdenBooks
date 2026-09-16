package dto

type CategoryResponse struct {
	ID          uint
	Name        string
	Description string
}

func NewCategoryResponse(ID uint, Name string, Description string) *CategoryResponse {
	return &CategoryResponse{ID: ID, Name: Name, Description: Description}
}

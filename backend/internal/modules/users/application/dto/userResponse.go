package dto

type UserResponse struct {
	ID        uint
	Email     string
	 
	Token     string //opcional, cuando llamas a getByID  no necesitas el token
	Name      string  //opcional
	Tel       string //opcional
	AddressID uint //opcional
	ImageURL  string //opcional
}

func NewUserResponse(id uint, email, token, name, tel string, addressID uint,ImageURL  string) *UserResponse {
	return &UserResponse{
		ID:        id,
		Email:     email,
		Token:     token,
		Name:      name,
		Tel:       tel,
		AddressID: addressID,
		ImageURL: ImageURL,
	}
}
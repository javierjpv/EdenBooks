package dto

type UserResponse struct {
	ID        uint
	Email     string
	 
	Token     string //opcional, cuando llamas a getByID  no necesitas el token
	Name      string  //opcional
	Tel       string //opcional
	AddressID uint //opcional
}

// func NewUserResponse()*UserResponse{

// }
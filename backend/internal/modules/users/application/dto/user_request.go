package dto

type UserRequest struct {
	Name      string `json:"name"` //SERA OPCIONAL EN LA SOLICITUD
	Email     string `json:"email"`
	Password  string `json:"password"`
	Tel       string `json:"tel"`
	AddressID *uint  `json:"addressID"`
	ImageURL  string  `json:"ImageURL"`
}

//NAME Y TEL SERAN OPCIONALES YA QUE ESTE DTO SE USARA TANTO EN REGISTER COMO
//EN LOGIN, PERO EN EL LOGIN SOLO SE NECESITARA EL EMAIL Y PASSWORD

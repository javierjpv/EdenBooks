package dto

type CarrierResponse struct {
	ID      uint
	Name    string 
	Contact string 
}

func NewCarrierResponse(ID uint,Name string, Contact string)*CarrierResponse{
	return &CarrierResponse{ID:ID,Name: Name,Contact: Contact}
}
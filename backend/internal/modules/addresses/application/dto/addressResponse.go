package dto

type AddressResponse struct {
	ID         uint
	City       string 
	Province   string 
	PostalCode string 
	Country    string 
	Street     string 
	Number     int    
}

func NewAddressResponse(ID uint,city string, province string, postalCode string, country string, street string, number int) *AddressResponse {
	return &AddressResponse{ID:ID,City: city, Province: province, PostalCode: postalCode, Country: country, Street: street, Number: number}
}

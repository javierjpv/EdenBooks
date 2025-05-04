package dto

type AddressRequest struct {
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
	Street     string `json:"street"`
	Number     int    `json:"number"`
}

func NewAddressRequest(city string, province string, postalCode string, country string, street string, number int) AddressRequest {
	return AddressRequest{City: city, Province: province, PostalCode: postalCode, Country: country, Street: street, Number: number}
}

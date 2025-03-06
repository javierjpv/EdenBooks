package dto



type AddressDTO struct{
	City string        `json:"city"`
    Province string    `json:"province"`
    PostalCode string  `json:"postalCode"`
	Country string     `json:"country"`
	Street     string  `json:"street"`
	Number     int     `json:"number"`
}
func NewAddressDTO(city string, province string, postalCode string, country string, street string,number int) *AddressDTO {
	return &AddressDTO{City: city, Province: province, PostalCode: postalCode, Country: country,Street: street,Number: number}
}
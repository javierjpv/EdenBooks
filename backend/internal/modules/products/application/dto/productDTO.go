package dto

type ProductDTO struct{
	Name string          `json:"Name"`
	Description string   `json:"Description"`
	Price float64        `json:"Price"`
	// OrderID uint         `json:"OrderID"` no se permitira crear un producto con un orderID ya que el order se rellenara en cuando se haga el pedido por lo tanto ese campo se actualizara
	CategoryID uint      `json:"CategoryID"`
	UserID uint          `json:"UserID"`
	ImageURL string      `json:"ImageURL"`
}
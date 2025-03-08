package usecases

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/services"
)

var (
    ErrMissingFields = fmt.Errorf("all fields are required")
    ErrInvalid     = fmt.Errorf("invalid ID")
)
type ProductUseCase struct {
	service services.ProductService	
}


func NewProductUseCase(service services.ProductService)*ProductUseCase{
	return &ProductUseCase{service: service}
}

func (u *ProductUseCase)CreateProduct(p dto.ProductDTO)error{
	if p.Name == "" || p.Description == "" ||p.ImageURL == ""{
		return ErrMissingFields
	}
	if p.Price<=0 {
		return fmt.Errorf("price can not be <= 0")
	}
	if  p.CategoryID==0 || p.UserID==0  {
		return ErrInvalid
	}
	return u.service.CreateProduct(p)
}

func (u *ProductUseCase)UpdateProduct(id uint,p dto.ProductDTO)error{
	if p.Name == "" || p.Description == "" || p.ImageURL == ""{
		return ErrMissingFields
	}
	if p.Price<=0 {
		return fmt.Errorf("price can not be <= 0")
	}
	if  p.CategoryID==0 || p.UserID==0  {
		return ErrInvalid
	}
	return u.service.UpdateProduct(id,p)
}

func (u *ProductUseCase)DeleteProduct(id uint)error{
	if  id==0  {
		return ErrInvalid
	}
	return u.service.DeleteProduct(id)
}

// func (u *ProductUseCase)GetProductByID(id uint)(*entities.Product,error){
// 	if  id==0  {
// 		return nil,ErrInvalid
// 	}
// 	return u.service.GetProductByID(id)
// }

func (u *ProductUseCase)GetProductByID(id uint,userID uint)(*entities.ProductWithFavoriteStatus, error){
	if  id==0  {
		return nil,ErrInvalid
	}
	return u.service.GetProductByIDWithFavorite(id,userID)
}
func (u *ProductUseCase)AddToFavorites(userID uint,productID uint)error{
	if  userID==0  {
		return ErrInvalid
	}
	if  userID==0  {
		return ErrInvalid
	}
	return u.service.AddToFavorites(userID,productID)
}

func (u *ProductUseCase)RemoveFromFavorites(userID uint,productID uint)error{
	if  userID==0  {
		return ErrInvalid
	}
	if  userID==0  {
		return ErrInvalid
	}
	return u.service.RemoveFromFavorites(userID,productID)
}
func (u *ProductUseCase)GetFavorites(userID uint) ([]entities.Product, error){
	if  userID==0  {
		return nil,ErrInvalid
	}
	return u.service.GetFavorites(userID)
}
func (u *ProductUseCase) GetFilteredProducts(filters map[string]string) ([]entities.Product, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"price": true, "name": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetFilteredProducts(filters)
}
func (u *ProductUseCase) GetProductsWithFavorites(userID uint, filters map[string]string) ([]entities.ProductWithFavoriteStatus, error) {
	// Validar el orden si está presente
	if order, exists := filters["order"]; exists {
		if order != "asc" && order != "desc" {
			filters["order"] = "asc" // Valor por defecto
		}
	}

	// Validar columna de orden si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		validSortColumns := map[string]bool{"price": true, "name": true}
		if !validSortColumns[sortBy] {
			delete(filters, "sort_by") // Eliminar si no es válido
		}
	}

	return u.service.GetProductsWithFavorites(userID,filters)
}
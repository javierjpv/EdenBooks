package services

import (
	"fmt"
	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/repositories"
	eventBusService "github.com/javierjpv/edenBooks/internal/shared/domain/services"
	// orderService "github.com/javierjpv/edenBooks/internal/modules/orders/domain/services"
)




type ProductService struct{
	repo repositories.ProductRepository
	eventBusService eventBusService.EventBus
	// orderService orderService.OrderService
}

func NewProductService(repo repositories.ProductRepository,eventBusService eventBusService.EventBus)*ProductService{
	return &ProductService{repo: repo,eventBusService: eventBusService}
}

func(s * ProductService)CreateProduct(p dto.ProductDTO)error{
	product:=entities.NewProduct(p.Name,p.Description,p.Price,p.CategoryID,p.UserID,p.ImageURL)//se creara un user sin order id ya q esta se llevara a cabo mas adelante
    return s.repo.CreateProduct(product)
}

func(s * ProductService)UpdateProduct(id uint, p dto.ProductDTO)error{
	product,err:=s.repo.GetProductByID(id)
    if err!=nil{
		return err
	}
	product.Name=p.Name
	product.Description=p.Description
	product.Price=p.Price
	product.CategoryID=p.CategoryID
	product.UserID=p.UserID
	product.ImageURL=p.ImageURL
	return s.repo.UpdateProduct(product)	
}
func(s * ProductService)AddOrderIDToProducts(orderID uint,productsIDs []uint)error{
	for _, productID := range productsIDs {
		product,err:=s.repo.GetProductByID(productID)
		if err!=nil{
			return fmt.Errorf("product does not exist")
		}
		if product.Sold {
			return fmt.Errorf("product has already been sold")
		}
	}
	for _, productsID := range productsIDs {
		product,err:=s.repo.GetProductByID(productsID)
		if err!=nil{
			return fmt.Errorf("product does not exist")
		}
		product.OrderID=&orderID
		product.Sold=true
		if err:=s.repo.UpdateProduct(product);err!=nil{
			return  fmt.Errorf("product orderID can not be updated")
		}
	}
		// Publicar evento en el Bus
		eventData := map[string]interface{}{
			"content": fmt.Sprintf("Se ha creado un pedido con el id: %d",orderID),
			"seen":false,
			"userID":uint(1), //cambiar en el futuro para que sea dinamico
		}
		fmt.Println("📢 Publicando evento 'order.created' con datos:", eventData)
		s.eventBusService.Publish("order.created", eventData)

	return nil
}

func(s * ProductService)DeleteProduct(id uint)error{
	if _,err:= s.repo.GetProductByID(id);err!=nil{
		return err
	}
	return s.repo.DeleteProduct(id)
}
func(s * ProductService)GetProductByID(id uint)(*entities.Product,error){
	return s.repo.GetProductByID(id)
}
func(s * ProductService)AddToFavorites(userID uint, productID uint)error{
	if _,err:= s.repo.GetProductByID(productID);err!=nil{
		return err
	}
	return s.repo.AddToFavorites(userID,productID)
}
func(s * ProductService)RemoveFromFavorites(userID uint, productID uint) error{
	if _,err:= s.repo.GetProductByID(productID);err!=nil{
		return err
	}
 return s.repo.RemoveFromFavorites(userID,productID)
}

func (s *ProductService) GetFavorites(userID uint) ([]entities.Product, error){
	return s.repo.GetFavorites(userID)
}

func (s *ProductService) GetFilteredProducts(filters map[string]string) ([]entities.Product, error) {
	return s.repo.GetFilteredProducts(filters)
}
func (s *ProductService)GetProductsWithFavorites(userID uint, filters map[string]string) ([]entities.ProductWithFavoriteStatus, error){
	return s.repo.GetProductsWithFavorites(userID,filters)
}

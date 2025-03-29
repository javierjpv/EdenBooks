package repositories

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/products/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *entities.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil

}
func (r *ProductRepository) UpdateProduct(product *entities.Product) error {
	err := r.db.Save(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	err := r.db.Delete(&entities.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetProductByID(id uint) (*entities.Product, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("error fetching product: %w", err)
	}
	return &product, nil
}
func (r *ProductRepository) GetProductByIDWithFavorite(id uint, userID uint) (*dto.ProductResponse, error) {
	var product entities.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("error fetching product: %w", err)
	}

	// Verificar si el producto está en la lista de favoritos del usuario
	var count int64
	err = r.db.Table("user_favourites").
		Where("user_id = ? AND product_id = ?", userID, id).
		Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("error checking favorite status: %w", err)
	}

	// Retornar el producto con el estado de favorito
	return &dto.ProductResponse{
		Product:    product,
		IsFavorite: count > 0, // Si count > 0, es favorito
	}, nil
}
func (r *ProductRepository) AddToFavorites(userID uint, productID uint) error {
	product := &entities.Product{}

	if err := r.db.First(product, productID).Error; err != nil {
		fmt.Println("Error: Producto con ID", productID, "no encontrado en la BD")
		return fmt.Errorf("product with ID %d not found", productID)
	}
	// Agregar a favoritos
	if err := r.db.Model(product).Association("FavoritedBy").Append(&userEntities.User{Model: gorm.Model{ID: userID}}); err != nil {
		fmt.Println("Error al agregar a favoritos:", err)
		return err
	}

	return nil
}
func (r *ProductRepository) GetFavorites(userID uint) ([]entities.Product, error) {
	var products []entities.Product

	// Realizamos el JOIN
	err := r.db.
		Joins("JOIN user_favourites uf ON uf.product_id = products.id").
		Where("uf.user_id = ?", userID).
		Find(&products).Error

	if err != nil {
		fmt.Println("Error al obtener los productos favoritos:", err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) RemoveFromFavorites(userID uint, productID uint) error {
	product := &entities.Product{}

	if err := r.db.First(product, productID).Error; err != nil {
		fmt.Println("product not found")
		return fmt.Errorf("product with ID %d not found", productID)
	}
	// Verificar si el usuario tiene este producto en favoritos
	var count int64
	if err := r.db.Table("user_favourites").Where("user_id = ? AND product_id = ?", userID, productID).Count(&count).Error; err != nil {
		fmt.Println("error buscando si esta en favoritos")
		return err
	}

	// Si no existe la relación, no intentamos borrar
	if count == 0 {
		fmt.Println("user does not have product in favorites")
		return fmt.Errorf("user with ID %d does not have product with ID %d in favorites", userID, productID)
	}

	if err := r.db.Model(product).Association("FavoritedBy").Delete(&userEntities.User{Model: gorm.Model{ID: userID}}); err != nil {
		return err
	}

	return nil
}
func (r *ProductRepository) GetFilteredProducts(filters map[string]string) ([]entities.Product, error) {
	var products []entities.Product
	query := r.db

	for key, value := range filters {
		switch key {
		case "name":
			query = query.Where("name ILIKE ?", "%"+value+"%")
		case "description":
			query = query.Where("description ILIKE ?", "%"+value+"%")
		case "min_price":
			query = query.Where("price >= ?", value)
		case "max_price":
			query = query.Where("price <= ?", value)
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "user_id":
			query = query.Where("user_id = ?", value)
		case "order_id":
			query = query.Where("order_id = ?", value)
		}
	}

	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := filters["order"]
		query = query.Order(sortBy + " " + order)
	}

	// Aplicar paginación si está presente
	limitInt := 50 // Límite por defecto
	if limit, exists := filters["limit"]; exists {
		parsedLimit, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = parsedLimit
		}
	}
	query = query.Limit(limitInt)

	if page, exists := filters["page"]; exists {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			query = query.Offset((pageInt - 1) * limitInt)
		}
	}

	// Ejecutar la consulta
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// Función para obtener los productos y marcar si están en favoritosGetProductsWithFavorites
func (r *ProductRepository) GetProductsWithFavorites(userID uint, filters map[string]string) ([]dto.ProductResponse, error) {
	var products []entities.Product
	var favoriteProductIDs []uint

	query := r.db.Model(&entities.Product{})

	// Aplica los filtros, paginación, etc.
	for key, value := range filters {
		switch key {
		case "name":
			query = query.Where("name ILIKE ?", "%"+value+"%")
		case "description":
			query = query.Where("description ILIKE ?", "%"+value+"%")
		case "min_price":
			query = query.Where("price >= ?", value)
		case "max_price":
			query = query.Where("price <= ?", value)
		case "category_id":
			query = query.Where("category_id = ?", value)
		case "exclude_user":
			query = query.Where("user_id != ?", value)
		case "user_id":
			query = query.Where("user_id = ?", value)
		case "order_id":
			query = query.Where("order_id = ?", value)
		}
	}
	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := "asc"
		if val, ok := filters["order"]; ok && (val == "asc" || val == "desc") {
			order = val
		}
		query = query.Order(sortBy + " " + order)
	}
	// Paginación
	if page, exists := filters["page"]; exists {
		if pageInt, err := strconv.Atoi(page); err == nil && pageInt > 0 {
			query = query.Offset((pageInt - 1) * 20).Limit(20)
		} else {
			query = query.Limit(20) // Valor por defecto
		}
	} else {
		query = query.Limit(20) // Valor por defecto si no se envía "page"
	}

	// Obtener todos los productos
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	// Si no hay productos, devolver lista vacía sin consultar favoritos
	if len(products) == 0 {
		return []dto.ProductResponse{}, nil
	}
	// Obtener los productos favoritos del usuario (IDs)
	if err := r.db.Table("user_favourites").
		Where("user_id = ?", userID).
		Pluck("product_id", &favoriteProductIDs).Error; err != nil {
		return nil, err
	}

	// Crear un slice de ProductWithFavoriteStatus para devolver
	productsWithStatus := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		productsWithStatus[i] = dto.ProductResponse{
			Product:    product,
			IsFavorite: contains(favoriteProductIDs, product.ID),
		}
	}

	return productsWithStatus, nil
}

// Helper para verificar si un ID está en la lista de favoritos
func contains(slice []uint, item uint) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

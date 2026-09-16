package repositories

import (
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/carriers/domain/entities"
	"gorm.io/gorm"
)

type CarrierRepository struct {
	db *gorm.DB //Este tipo de objeto es el que te devuelve al usar gorm.open (conexion a la base de datos)
}

func NewCarrierRepository(db *gorm.DB) *CarrierRepository {
	return &CarrierRepository{db: db}
}

func (r *CarrierRepository) CreateCarrier(c *entities.Carrier) error {
	if err := r.db.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (r *CarrierRepository) UpdateCarrier(c *entities.Carrier) error {
	result := r.db.Model(&entities.Carrier{}).Where("id = ?", c.ID).Updates(c)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *CarrierRepository) DeleteCarrier(id uint) error {
    result := r.db.Delete(&entities.Carrier{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}

func (r *CarrierRepository) GetCarrierByID(id uint) (*entities.Carrier, error) {
	var carrier entities.Carrier //aqui es necesario crearla ya que se guardara ahi mismo el resultado
	err := r.db.First(&carrier, id).Error
	if err != nil {
		return nil, err
	}
	return &carrier, nil
}

func (r *CarrierRepository) GetFilteredCarrieres(filters map[string]string) ([]entities.Carrier, error) {
	var carrieres []entities.Carrier
	query := r.db

	// Aplicar dinámicamente los filtros
	for key, value := range filters {
		switch key {
		case "start_date":
			query = query.Where("created_at >= ?", value)
		case "end_date":
			query = query.Where("created_at <= ?", value)
		case "name":
			query = query.Where("name = ?", value)
		case "contact":
			query = query.Where("contact = ?", value)
		}
	}

	// Aplicar ordenamiento si está presente
	if sortBy, exists := filters["sort_by"]; exists {
		order := filters["order"]
		query = query.Order(sortBy + " " + order)
	}

	// Aplicar paginación si está presente
	if page, exists := filters["page"]; exists {
		limit := filters["limit"]
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		query = query.Offset((pageInt - 1) * limitInt).Limit(limitInt)
	}

	// Ejecutar la consulta
	if err := query.Find(&carrieres).Error; err != nil {
		return nil, err
	}

	return carrieres, nil
}

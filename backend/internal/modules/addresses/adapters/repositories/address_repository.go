package repositories

import (
	// "fmt"
	"errors"
	"strconv"

	"github.com/javierjpv/edenBooks/internal/modules/addresses/domain/entities"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB //Este tipo de objeto es el que te devuelve al usar gorm.open (conexion a la base de datos)
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (r *AddressRepository) CreateAddress(address *entities.Address) (*entities.Address, error) {
	existingAddress, exist, err := r.CheckExistingAddress(address)
	if err != nil {
		return nil, err
	}
	if exist {
		return existingAddress, nil
	}
	result := r.db.Create(address)
	if result.Error != nil {
		return nil, err
	}
	return address, nil
}

func (r *AddressRepository) UpdateAddress(address *entities.Address) error {
	result := r.db.Model(&entities.Address{}).Where("id = ?", address.ID).Updates(address)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *AddressRepository) DeleteAddress(id uint) error {
	result := r.db.Delete(&entities.Address{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *AddressRepository) GetAddressByID(id uint) (*entities.Address, error) {
	var address entities.Address //aqui es necesario crearla ya que se guardara ahi mismo el resultado
	err := r.db.First(&address, id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}
func (r *AddressRepository) CheckExistingAddress(address *entities.Address) (*entities.Address, bool, error) {
	var existingAddress entities.Address

	err := r.db.Where(
		"country = ? AND province = ? AND city = ? AND postal_code = ? AND number = ? AND street = ?",
		address.Country, address.Province, address.City, address.PostalCode, address.Number, address.Street,
	).First(&existingAddress).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &existingAddress, true, nil
}

func (r *AddressRepository) GetFilteredAddresses(filters map[string]string) ([]entities.Address, error) {
	var addresses []entities.Address
	query := r.db

	// Aplicar dinámicamente los filtros
	for key, value := range filters {
		switch key {
		case "start_date":
			query = query.Where("created_at >= ?", value)
		case "end_date":
			query = query.Where("created_at <= ?", value)
		case "city":
			query = query.Where("city = ?", value)
		case "province":
			query = query.Where("province = ?", value)
		case "postal_code":
			query = query.Where("postal_code = ?", value)
		case "country":
			query = query.Where("country = ?", value)
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
	if err := query.Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

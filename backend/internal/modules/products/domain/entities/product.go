package entities

import (
	// reviewEntities "github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	userEntities "github.com/javierjpv/edenBooks/internal/modules/users/domain/entities"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string  `gorm:"not null"`
	Description   string  `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	OrderID       *uint //relacion 1:N un pedido tiene muchos productos y un producto soloe esta en un pedido
	//Order deberia ser opcional ya que al crear un producto no debe tener un order
	CategoryID uint                    `gorm:"not null"` //Relacion 1:N una categoria tiene muchos productos y un producto tiene una categoria
	UserID     uint                    `gorm:"not null"` //Relacion 1:N un produto tiene un usuario y un usuario tiene muchos productos
	ImageURL   string                  `gorm:"not null"`
	Sold       bool                    `gorm:"default:false"`
	//No realizaras el preload de reviews asi que puedes comentarlo o dejarlo
	// Reviews    []reviewEntities.Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //Relacion 1:N
	FavoritedBy []userEntities.User `gorm:"many2many:user_favourites;"`
}

func NewProduct(name string, description string, price float64,  categoryID uint, userID uint, ImageURL string) *Product {
	return &Product{Name: name, Description: description, Price: price, CategoryID: categoryID, UserID: userID, ImageURL: ImageURL}
}

// Cuando incluyes gorm.Model en tu estructura, automáticamente obtienes los siguientes campos:

// ID (tipo uint): Es el identificador único de la entidad. En bases de datos como SQLite, este campo se autoincrementa por defecto.

// CreatedAt (tipo time.Time): Es la marca de tiempo que indica cuándo fue creado el registro.

// UpdatedAt (tipo time.Time): Es la marca de tiempo que indica cuándo fue la última vez que se actualizó el registro.

// DeletedAt (tipo *time.Time): Es la marca de tiempo que indica cuándo se eliminó el registro (esto solo se utiliza si activas el "soft delete", es decir, cuando GORM marca los registros como eliminados sin borrarlos físicamente).

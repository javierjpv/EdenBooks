package entities

import (
	orderEntities "github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	"time"
)

type Carrier struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"not null;unique"` // Ej: "DHL", "FedEx", "UPS"
	Contact   string
	Orders    []orderEntities.Order
}

func NewCarrier(name, contact string) *Carrier {
	return &Carrier{Name: name, Contact: contact}
}

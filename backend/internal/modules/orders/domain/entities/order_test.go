
package entities_test

import (
	"testing"

	"github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
)

func TestNewOrder(t *testing.T) {
	// 1. Definición de la estructura de casos de prueba (Table-Driven)
	tests := []struct {
		name          string
		state         string
		userID        uint
		addressID     uint
		carrierID     uint
		transactionID uint
	}{
		{
			name:          "debe crear una orden válida correctamente",
			state:         "PENDING",
			userID:        1,
			addressID:     10,
			carrierID:     5,
			transactionID: 99,
		},
		{
			name:          "debe permitir crear una orden con estado PAGADO",
			state:         "PAID",
			userID:        2,
			addressID:     12,
			carrierID:     3,
			transactionID: 100,
		},
	}

	// 2. Ejecución iterativa de casos
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := entities.NewOrder(
				tt.state,
				tt.userID,
				tt.addressID,
				tt.carrierID,
				tt.transactionID,
			)

			// 3. Aseveraciones (Assertions)
			if order == nil {
				t.Fatal("se esperaba una instancia de Order, se obtuvo nil")
			}
			if order.State != tt.state {
				t.Errorf("se esperaba el estado %s, se obtuvo %s", tt.state, order.State)
			}
			if order.UserID != tt.userID {
				t.Errorf("se esperaba UserID %d, se obtuvo %d", tt.userID, order.UserID)
			}
			if order.AddressID != tt.addressID {
				t.Errorf("se esperaba AddressID %d, se obtuvo %d", tt.addressID, order.AddressID)
			}
		})
	}
}
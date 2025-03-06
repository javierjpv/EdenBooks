package entities

import (
	messageEntities "github.com/javierjpv/edenBooks/internal/modules/messages/domain/entities"
	notificationEntities "github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
	orderEntities "github.com/javierjpv/edenBooks/internal/modules/orders/domain/entities"
	productEntities "github.com/javierjpv/edenBooks/internal/modules/products/domain/entities"
	reviewEntities "github.com/javierjpv/edenBooks/internal/modules/reviews/domain/entities"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Name string  `gorm:"default:''"`
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`
	// Role string //poner user por defecto

	Tel string   `gorm:"default:''"`
	AddressID *uint  `gorm:"default:NULL"`
	//En Gorm en las relaciones 1:N la clave foranea ira hacia la tabla N
	//En estos casos donde habra un array al haber colocado ya la clave foranea donde corresponde Gorm podra relacionarlos y entonces cuando el usuario tenga por ejemplo orders
	//relacionadas con un userId se podra acceder a este array mediante  Preload
	//Los slices no hace falta ponerlos como nullables ya que por defecto son slices vacios ([])
	Products []productEntities.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N
	Orders []orderEntities.Order   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N
    Reviews []reviewEntities.Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N
    Notifications []notificationEntities.Notification `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relacion 1:N
	// Chats [] chatEntities.Chat       `gorm:"many2many:user_chats;"` // Relacion N:M   //Debes arreglarlo para que se pueda acceder desde los chat a los users 
	//Actualmente solo los users pueden acceder a sus chats y no es bidireccional
	//La tabla intermedia entre user y chat es creada  igulmente
	// Relación 1:N (el usuario puede enviar muchos mensajes)
	SentMessages []messageEntities.Message `gorm:"foreignkey:SenderID"` // Mensajes enviados por el usuario
	// Relación 1:N (el usuario puede recibir muchos mensajes)
	ReceivedMessages []messageEntities.Message `gorm:"foreignkey:ReceiverID"` // Mensajes recibidos por el usuario

}

func NewUser(email string, password string)*User{
    
    return &User{Email: email,Password:password}
}



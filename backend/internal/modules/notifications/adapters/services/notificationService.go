package services

import (
	"fmt"

	"github.com/javierjpv/edenBooks/internal/modules/notifications/application/dto"
	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/entities"
	"github.com/javierjpv/edenBooks/internal/modules/notifications/domain/repositories"
	eventBusService "github.com/javierjpv/edenBooks/internal/shared/domain/services"
)

type NotificationService struct{
	repo repositories.NotificationRepository
	eventBusService eventBusService.EventBus
}

func NewNotificationService(repo repositories.NotificationRepository,eventBusService eventBusService.EventBus)*NotificationService{
	return &NotificationService{repo: repo,eventBusService: eventBusService}
}

func(s * NotificationService)CreateNotification(n dto.NotificationDTO)error{
	notification:=entities.NewNotification(n.Content,n.Seen,n.UserID)
    return s.repo.CreateNotification(notification)
}

func(s * NotificationService)UpdateNotification(id uint, n dto.NotificationDTO)error{
	notification,err:=s.repo.GetNotificationByID(id)
    if err!=nil{
		return err
	}
    notification.Content=n.Content
	notification.Seen=n.Seen
	notification.UserID=n.UserID
	return s.repo.UpdateNotification(notification)	
}
func(s * NotificationService)DeleteNotification(id uint)error{
	if _,err:=s.repo.GetNotificationByID(id);err!=nil{
		return err
	}
	return s.repo.DeleteNotification(id)
}
func(s * NotificationService)GetNotificationByID(id uint)(*entities.Notification,error){
	return s.repo.GetNotificationByID(id)
}
func (s *NotificationService) GetFilteredNotifications(filters map[string]string) ([]entities.Notification, error) {
	return s.repo.GetFilteredNotifications(filters)
}
// Suscribir al evento
func (s *NotificationService) ListenOrderCreated() {
	err := s.eventBusService.Subscribe("order.created", func(data interface{}) {
		fmt.Println("Evento recibido en NotificationService")
		eventData, ok := data.(map[string]interface{})
		if !ok {
			fmt.Println("Error al procesar el evento")
			return
		}
		// Convertir los valores a los tipos correctos
				content, ok := eventData["content"].(string)
		if !ok {
			fmt.Println("Error al convertir 'content' a string")
			return
		}
		
		seen, ok := eventData["seen"].(bool)
		if !ok {
			fmt.Println("Error al convertir 'seen' a bool")
			return
		}

		userID, ok := eventData["userID"].(uint) // Ahora verificamos directamente si es uint
		if !ok {
			fmt.Println("Error al convertir 'userID' a uint")
			return
		}
					
		notification:=entities.NewNotification(content,seen,userID)
		if err:=s.repo.CreateNotification(notification);err!=nil {
			fmt.Println("Error al crear la notificacion") // habra un mecanismo de reintento
			return
		}
		fmt.Printf("🔔 Notificación: Pedido %v\n", content)
		fmt.Printf("Estado: Visto: %v, UsuarioID: %v\n", seen, userID)
		
		
	})
	
	if err != nil {
		fmt.Println("Error al suscribirse al evento:", err)
	}
}
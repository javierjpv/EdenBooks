package handlers

import (
	"encoding/json"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// func (h *Hub) Run() {
// 	for {
// 		select {
// 		case client := <-h.register:
// 			h.clients[client] = true
// 		case client := <-h.unregister:
// 			if _, ok := h.clients[client]; ok {
// 				delete(h.clients, client)
// 				close(client.send)
// 			}
// 		case message := <-h.broadcast:
// 			for client := range h.clients {
// 				select {
// 				case client.send <- message:
// 				default:
// 					close(client.send)
// 					delete(h.clients, client)
// 				}
// 			}
// 		}
// 	}
// }


func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case message := <-h.broadcast:
            // Parsear el mensaje para obtener los IDs de remitente y destinatario
            var messageData map[string]interface{}
            if err := json.Unmarshal(message, &messageData); err != nil {
                log.Printf("Error al parsear el mensaje: %v", err)
                continue
            }
            
            // Extraer el ID del destinatario
            receiverIDFloat, ok := messageData["ReceiverID"].(float64)
            if !ok {
                log.Println("receiverID inválido o faltante en el mensaje")
                continue
            }
            receiverID := uint(receiverIDFloat)
            
            // Extraer el ID del remitente
            senderIDFloat, ok := messageData["SenderID"].(float64)
            if !ok {
                log.Println("senderID inválido o faltante en el mensaje")
                continue
            }
            senderID := uint(senderIDFloat)
            
            // Enviar el mensaje solo al remitente y al destinatario
            for client := range h.clients {
                // Enviar el mensaje si el cliente es el remitente o el destinatario
                if client.userID == senderID || client.userID == receiverID {
                    select {
                    case client.send <- message:
                    default:
                        close(client.send)
                        delete(h.clients, client)
                    }
                }
            }
        }
    }
}
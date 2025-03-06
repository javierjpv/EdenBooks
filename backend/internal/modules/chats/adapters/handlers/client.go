package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/javierjpv/edenBooks/internal/modules/messages/application/dto"
	messageUsecase "github.com/javierjpv/edenBooks/internal/modules/messages/application/useCases"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		log.Println("Origin:", r.Header.Get("Origin"))
		allowedOrigins := map[string]bool{
			"http://localhost:5173": true,
		}
		return allowedOrigins[r.Header.Get("Origin")]
	},
}

// Client is a middleman between the websocket connection and the hub.

// type Client struct {
// 	hub *Hub

// 	// The websocket connection.
// 	conn *websocket.Conn

// 	// Buffered channel of outbound messages.
// 	send chan []byte

// }

type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	userID uint

}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump(messageUsecase *messageUsecase.MessageUseCase) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		var messageDto dto.MessageDTO
		err = json.Unmarshal(message, &messageDto)
		if err != nil {
			log.Printf("Error al deserializar el mensaje: %v", err)
			continue
		}
        messageUsecase.CreateMessage(messageDto)

		
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request,messageUseCase *messageUsecase.MessageUseCase) {
	//leer el token aqui antes de actualizar r
    // Obtener el token del query parameter
    tokenString := r.URL.Query().Get("token")
    if tokenString == "" {
        // Si no está en los query parameters, intentar obtenerlo del header
        authHeader := r.Header.Get("Authorization")
        if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
            tokenString = strings.TrimPrefix(authHeader, "Bearer ")
        }
    }
    
    // Validar el token y obtener el ID del usuario
    userID, err := ExtractUserIDFromToken(tokenString)
    if err != nil {
        log.Printf("Error al validar el token: %v", err)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256),userID: userID}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump(messageUseCase)
}

func ExtractUserIDFromToken(tokenString string) (uint, error) {
    // Usar la misma clave secreta que se usa para firmar
    var secretKey = []byte("supersecret")
    
    // Parsear el token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verificar que el método de firma sea el correcto
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
        }
        return secretKey, nil
    })
    
    if err != nil {
        return 0, fmt.Errorf("error al validar el token: %v", err)
    }
    
    // Verificar que el token sea válido
    if !token.Valid {
        return 0, errors.New("token inválido")
    }
    
    // Extraer los claims (información) del token
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("no se pudieron extraer los claims del token")
    }
    
    // Extraer el ID del usuario
    userIDFloat, ok := claims["ID"].(float64)
    if !ok {
        return 0, errors.New("ID de usuario no encontrado en el token")
    }
    
    // Convertir a uint
    userID := uint(userIDFloat)
    
    return userID, nil
}
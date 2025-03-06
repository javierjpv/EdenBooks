package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	messageUsecase "github.com/javierjpv/edenBooks/internal/modules/messages/application/useCases"
	chatHandler "github.com/javierjpv/edenBooks/internal/modules/chats/adapters/handlers"
)

var addr = flag.String("addr", ":8080", "http service address")

func Run(messageUsecase *messageUsecase.MessageUseCase) {
	flag.Parse()
	hub := chatHandler.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Configurar CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		chatHandler.ServeWs(hub, w, r,messageUsecase)
	})
	fmt.Printf("Server listening on ws://localhost%s/ws\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pascallin/go-communication/providers"
	"github.com/pascallin/go-communication/services"
)

var upgrader = websocket.Upgrader{
	// CORS
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Communication(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)

			break
		}
		log.Printf("recv: %s", message)
		// debug
		providers.DispatchToProvider(1, providers.Payload{
			KeyWords: services.TokenizeString(string(message)),
		})
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

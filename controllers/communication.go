package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pascallin/go-communication/models"
	"github.com/pascallin/go-communication/protocol"
	"github.com/pascallin/go-communication/repositories"
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
		// insert message to mongo
		data := protocol.Decode(message)
		repositories.InsertMessage(&models.Message{
			Author:  "pascal",
			Message: data.Message,
		})
		// debug
		// providers.DispatchToProvider(1, providers.Payload{
		// 	KeyWords: services.TokenizeString(string(message)),
		// })
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

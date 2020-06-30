package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	Controllers "github.com/pascallin/go-communication/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var addr = flag.String("addr", "localhost:"+os.Getenv("PORT"), "http service address")

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", Controllers.Echo)
	http.HandleFunc("/messages", Controllers.GetMessageList)
	http.HandleFunc("/messengers", Controllers.GetMessengerList)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

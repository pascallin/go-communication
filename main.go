package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pascallin/go-communication/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var addr = flag.String("addr", "localhost:"+os.Getenv("PORT"), "http service address")

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/communication", controllers.Communication)
	http.HandleFunc("/messages", controllers.GetMessageList)
	http.HandleFunc("/messengers", controllers.GetMessengerList)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

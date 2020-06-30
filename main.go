package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pascallin/go-communication/controllers"
	"github.com/pascallin/go-communication/databases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// server address flag
	var addr = flag.String("addr", "localhost:"+os.Getenv("PORT"), "http service address")

	// connect mongodb
	mongo, err := databases.NewMongoDatabase()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/communication", controllers.Communication)
	http.HandleFunc("/messages", controllers.GetMessageList)
	http.HandleFunc("/messengers", controllers.GetMessengerList)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

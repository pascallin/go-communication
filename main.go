package main

import (
	"flag"
	"log"
	"net/http"

	Controllers "github.com/pascallin/go-communication/controllers"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", Controllers.Echo)
	http.HandleFunc("/messages", Controllers.GetMessageList)
	http.HandleFunc("/messengers", Controllers.GetMessengerList)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

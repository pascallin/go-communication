package main

import (
	"flag"
	"github.com/pascallin/go-communication/internal/message"
	"github.com/pascallin/go-communication/internal/messenger"
	"github.com/pascallin/go-communication/internal/pkg/databases"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func main() {
	flag.Parse()
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	defer glog.Flush()

	err := godotenv.Load()
	if err != nil {
		glog.Fatal("Error loading .env file")
	}

	var addr = "localhost:"+os.Getenv("PORT")

	// connect mongodb
	mongo, err := databases.NewMongoDatabase()
	if err != nil {
		panic(err)
	}
	defer mongo.Close()

	http.HandleFunc("/communication", message.Communication)
	http.HandleFunc("/messages", message.GetMessageList)
	http.HandleFunc("/messengers", messenger.GetMessengerList)

	glog.Infof("server listening %s", addr)
	if err = http.ListenAndServe(addr, nil); err != nil {
		glog.Fatal(err)
	}
}

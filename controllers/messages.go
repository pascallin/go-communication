package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pascallin/go-communication/repositories"
)

type Message struct {
	Id        int64  `json:"id"`
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type MessageData struct {
	TempMessages []Message `json:"tempMessages"`
}

func GetMessageList(w http.ResponseWriter, r *http.Request) {
	// file, err := ioutil.ReadFile("datasources/message.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	result := repositories.GetMessages(1, 20)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

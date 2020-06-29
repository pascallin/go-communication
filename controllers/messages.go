package Controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	file, err := ioutil.ReadFile("datasources/message.json")
	if err != nil {
		fmt.Println(err)
	}
	// data := MessageData{}
	// err := json.Unmarshal([]byte(file), &data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(file))
}

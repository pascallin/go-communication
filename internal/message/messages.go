package message

import (
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageData struct {
	TempMessages []Message `json:"tempMessages"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Author    string             `bson:"author" json:"author"`
	Message   string             `bson:"message" json:"message"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

func (t *Message) New() *Message {
	return &Message{
		ID:        primitive.NewObjectID(),
		Author:    t.Author,
		Message:   t.Message,
		Timestamp: t.Timestamp,
	}
}


func GetMessageList(w http.ResponseWriter, r *http.Request) {
	// file, err := ioutil.ReadFile("datasources/message.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	result := GetMessages(1, 20)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

package message

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pascallin/go-communication/internal/pkg/databases"
	"github.com/pascallin/go-communication/internal/pkg/protocol"
	"github.com/pascallin/go-communication/internal/pkg/tokenize"
)

var collectionName string = "messages"

func GetMessages(page int64, pageSize int64) []*Message {
	var results []*Message
	ctx := context.Background()

	// init condition
	condition := bson.D{}
	findOptions := options.Find()
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (page - 1))
	findOptions.SetSort(bson.D{{"timestamp", 1}})

	cur, err := databases.MongoDB.DB.Collection(collectionName).Find(ctx, condition, findOptions)
	if err != nil {
		return nil
	}
	// Close the cursor once finished
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var message Message
		err := cur.Decode(&message)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &message)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

func InsertMessage(m *Message) *Message {
	ctx := context.Background()

	m.ID = primitive.NewObjectID()
	m.Timestamp = time.Now()

	insertResult, err := databases.MongoDB.DB.Collection(collectionName).InsertOne(ctx, m)
	if err != nil {
		return nil
	}

	fmt.Println("Inserted message: ", insertResult.InsertedID)
	return m
}

func Communication(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Upgrader{
		// CORS
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}.Upgrade(w, r, nil)
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
		InsertMessage(&Message{
			Author:  "pascal",
			Message: data.Message,
		})
		// debug tokenize
		DispatchToProvider(1, Payload{
			KeyWords: tokenize.TokenizeString(string(message)),
		})
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
		InsertMessage(&Message{
			Author:  "system",
			Message: string(message),
		})
	}
}
package message

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pascallin/go-communication/internal/pkg/databases"
	"github.com/pascallin/go-communication/internal/pkg/protocol"
	"github.com/pascallin/go-communication/internal/pkg/tokenize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var collectionName string = "messages"

func GetMessages(page int64, pageSize int64) []*Message {
	var results []*Message
	ctx := context.Background()

	// init condition
	condition := bson.D{}
	findOptions := options.Find()
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(page * (pageSize - 1))
	findOptions.SetSort(bson.D{{"timestamp", 1}})

	cur, err := databases.MongoDB.DB.Collection(collectionName).Find(ctx, condition, findOptions)
	if err != nil {
		return nil
	}
	fmt.Printf("cur: %+v\n", cur)
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
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
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
		InsertMessage(&Message{
			Author:  "pascal",
			Message: data.Message,
		})
		// debug
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
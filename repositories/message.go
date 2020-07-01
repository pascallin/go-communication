package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pascallin/go-communication/databases"
	"github.com/pascallin/go-communication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "messages"

func GetMessages(page int64, pageSize int64) []*models.Message {
	var results []*models.Message
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
		var message models.Message
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

func InsertMessage(m *models.Message) *models.Message {
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

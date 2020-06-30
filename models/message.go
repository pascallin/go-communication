package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

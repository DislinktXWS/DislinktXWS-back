package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	FirstParticipator  string             `bson:"first_participator"`
	SecondParticipator string             `bson:"second_participator"`
	Messages           []Message          `bson:"messages"`
}

type Message struct {
	Author  string `bson:"author"`
	Content string `bson:"content"`
	Date    string `bson:"date"`
}

package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Content string             `bson:"content"`
	Date    string             `bson:"date"`
	User    string             `bson:"user"`
}

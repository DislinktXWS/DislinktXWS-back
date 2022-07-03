package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	From       User               `bson:"from"`
	To         User               `bson:"to"`
	Date       string             `bson:"date"`
	Content    string             `bson:"content"`
	IsReviewed bool               `bson:"isReviewed"`
}

type User struct {
	Username string `bson:"username"`
	Name     string `bson:"name"`
	Surname  string `bson:"surname"`
}

package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	PostId  string `bson:"postId"`
	User    string `bson:"user"`
	Content string `bson:"content"`
}

type Post struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Content  string             `bson:"content"`
	Links    []string           `bson:"links"`
	Date     string             `bson:"date"`
	User     string             `bson:"user"`
	Likes    []string           `bson:"likes"`
	Dislikes []string           `bson:"dislikes"`
	Comments []Comment          `bson:"comments"`
}

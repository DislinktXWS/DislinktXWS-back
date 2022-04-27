package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Insert(user *Post) error
	DeleteAll()
}

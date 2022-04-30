package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	Get(id primitive.ObjectID) (*Post, error)
	GetAll() ([]*Post, error)
	Insert(user *Post) error
	GetPostsByUser(user string) ([]*Post, error)
	LikePost(id primitive.ObjectID, username string)
	DislikePost(id primitive.ObjectID, username string)
	CommentPost(comment *Comment) error
	DeleteAll()
}

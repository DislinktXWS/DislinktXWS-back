package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostStore interface {
	Get(id primitive.ObjectID, ctx context.Context) (*Post, error)
	GetAll(ctx context.Context) ([]*Post, error)
	Insert(user *Post, ctx context.Context) error
	GetPostsByUser(user string, ctx context.Context) ([]*Post, error)
	LikePost(id primitive.ObjectID, username string, ctx context.Context)
	DislikePost(id primitive.ObjectID, username string, ctx context.Context)
	CommentPost(comment *Comment, ctx context.Context) error
	DeleteAll(ctx context.Context)
}

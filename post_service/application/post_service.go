package application

import (
	"context"
	"github.com/dislinktxws-back/post_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) Get(id primitive.ObjectID, ctx context.Context) (*domain.Post, error) {
	return service.store.Get(id, ctx)
}

func (service *PostService) GetAll(ctx context.Context) ([]*domain.Post, error) {
	return service.store.GetAll(ctx)
}

func (service *PostService) Insert(Post *domain.Post, ctx context.Context) error {
	return service.store.Insert(Post, ctx)
}

func (service *PostService) GetPostsByUser(user string, ctx context.Context) ([]*domain.Post, error) {
	return service.store.GetPostsByUser(user, ctx)
}

func (service *PostService) LikePost(id primitive.ObjectID, username string, ctx context.Context) {
	service.store.LikePost(id, username, ctx)
}

func (service *PostService) DislikePost(id primitive.ObjectID, username string, ctx context.Context) {
	service.store.DislikePost(id, username, ctx)
}

func (service *PostService) CommentPost(Comment *domain.Comment, ctx context.Context) error {
	return service.store.CommentPost(Comment, ctx)
}

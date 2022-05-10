package application

import (
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

func (service *PostService) Get(id primitive.ObjectID) (*domain.Post, error) {
	return service.store.Get(id)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) Insert(Post *domain.Post) error {
	return service.store.Insert(Post)
}

func (service *PostService) GetPostsByUser(user string) ([]*domain.Post, error) {
	return service.store.GetPostsByUser(user)
}

func (service *PostService) LikePost(id primitive.ObjectID, username string) {
	service.store.LikePost(id, username)
}

func (service *PostService) DislikePost(id primitive.ObjectID, username string) {
	service.store.DislikePost(id, username)
}

func (service *PostService) CommentPost(Comment *domain.Comment) error {
	return service.store.CommentPost(Comment)
}

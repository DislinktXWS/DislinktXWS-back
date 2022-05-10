package api

import (
	"github.com/dislinktxws-back/post_service/application"
)

type CreatePostCommandHandler struct {
	postService *application.PostService
}

func NewCreateUserCommandHandler(postService *application.PostService) (*CreatePostCommandHandler, error) {
	o := &CreatePostCommandHandler{
		postService: postService,
	}
	return o, nil
}

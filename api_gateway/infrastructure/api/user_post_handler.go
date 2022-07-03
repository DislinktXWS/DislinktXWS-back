package api

import (
	"context"
	"encoding/json"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	post_proto "github.com/dislinktxws-back/common/proto/post_service"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type UserPostHandler struct {
	userClientAddress string
	postClientAddress string
}

func NewUserPostHandler(userClientAddress, postClientAddress string) Handler {
	return &UserPostHandler{
		userClientAddress: userClientAddress,
		postClientAddress: postClientAddress,
	}
}

func (handler *UserPostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/posts/publicPosts", handler.GetPublicPosts)
	if err != nil {
		panic(err)
	}
}

func (handler *UserPostHandler) GetPublicPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)
	posts, err := postClient.GetAll(context.TODO(), &post_proto.GetAllRequest{})
	if err != nil {
		panic(err)
	}
	userClient := services.NewUserClient(handler.userClientAddress)
	users, err := userClient.GetPublicUsers(context.TODO(), &user_proto.GetPublicUsersRequest{})
	if err != nil {
		panic(err)
	}
	var publicPosts []*domain.Post
	for _, post := range posts.Posts {
		exists := false
		for _, user := range users.Users {
			if post.User == user.Id {
				exists = true
			}
		}
		if exists {
			publicPosts = append(publicPosts, mapNewPost(post))
		}
	}

	response, err := json.Marshal(publicPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

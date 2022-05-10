package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	connection_proto "github.com/dislinktxws-back/common/proto/connection_service"
	post_proto "github.com/dislinktxws-back/common/proto/post_service"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type UserFeedHandler struct {
	connectionClientAddress string
	postClientAddress       string
}

func NewUserFeedHandler(connectionClientAddress, postClientAddress string) Handler {
	return &UserFeedHandler{
		connectionClientAddress: connectionClientAddress,
		postClientAddress:       postClientAddress,
	}
}

func (handler *UserFeedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/feed/{id}", handler.GetUserFeed)
	if err != nil {
		panic(err)
	}
}

func (handler *UserFeedHandler) GetUserFeed(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Print("ID je -> ")
	fmt.Print(id)

	userIds, err := handler.getConnectionIds(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//sad sve postove za te id-eve
	posts := []domain.Post{}

	_ = handler.getPosts(&posts, userIds)

	response, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *UserFeedHandler) getConnectionIds(userId string) ([]string, error) {

	connectionsClient := services.NewConnectionClient(handler.connectionClientAddress)
	connections, err := connectionsClient.GetAll(context.TODO(), &connection_proto.GetAllConnectionsRequest{Id: userId})
	return connections.Ids, err
}

func (handler *UserFeedHandler) getPosts(posts *[]domain.Post, userIds []string) error {

	postClient := services.NewPostClient(handler.postClientAddress)

	for _, id := range userIds {
		userPosts, _ := postClient.GetPostsByUser(context.TODO(), &post_proto.GetPostsByUserRequest{User: id})
		for _, post := range userPosts.Posts {
			post := mapNewPost(post)
			*posts = append(*posts, *post)
		}
	}
	return nil
}

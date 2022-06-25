package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	connection_proto "github.com/dislinktxws-back/common/proto/connection_service"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type UserRecommendationsHandler struct {
	userClientAddress       string
	connectionClientAddress string
}

func NewUserRecommendationsHandler(userClientAddress, connectionClientAddress string) Handler {
	return &UserRecommendationsHandler{
		userClientAddress:       userClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserRecommendationsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/connections/recommendations/{id}", handler.GetUserRecommendations)
	if err != nil {
		panic(err)
	}
}

func (handler *UserRecommendationsHandler) GetUserRecommendations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Print("ID je -> ")
	fmt.Print(id)

	userIds, err := handler.getUserIds(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users := []domain.UserBasicInfo{}

	_ = handler.getUsers(&users, userIds)

	response, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *UserRecommendationsHandler) getUserIds(userId string) ([]string, error) {

	connectionsClient := services.NewConnectionClient(handler.connectionClientAddress)
	connections, err := connectionsClient.GetBlockedUsers(context.TODO(), &connection_proto.GetAllConnectionsRequest{Id: userId})
	return connections.Ids, err
}

func (handler *UserRecommendationsHandler) getUsers(users *[]domain.UserBasicInfo, userIds []string) error {

	userClient := services.NewUserClient(handler.userClientAddress)

	for _, id := range userIds {
		user, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: id})
		domainUser := mapNewUser(user.User)
		*users = append(*users, *domainUser)
	}
	return nil
}

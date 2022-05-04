package api

import (
	"context"
	"encoding/json"
	"fmt"
	"module/api_gateway/domain"
	"module/api_gateway/infrastructure/services"
	connection_proto "module/common/proto/connection_service"

	user_proto "module/common/proto/user_service"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type UserConnectionsHandler struct {
	userClientAddress       string
	connectionClientAddress string
}

func NewUserConnectionsHandler(userClientAddress, connectionClientAddress string) Handler {
	return &UserConnectionsHandler{
		userClientAddress:       userClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserConnectionsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/connections/{id}", handler.GetUserConnections)
	if err != nil {
		panic(err)
	}
}

//ovo je logika za slucaj da dobavljamo konekcije preko id, ako je nesto drugo onda jos jedan poziv user servisa

func (handler *UserConnectionsHandler) GetUserConnections(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

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

func (handler *UserConnectionsHandler) getUserIds(userId string) ([]string, error) {

	connectionsClient := services.NewConnectionClient(handler.connectionClientAddress)
	connections, err := connectionsClient.GetAll(context.TODO(), &connection_proto.GetAllConnectionsRequest{Id: userId})
	return connections.Ids, err
}

func (handler *UserConnectionsHandler) getUsers(users *[]domain.UserBasicInfo, userIds []string) error {

	userClient := services.NewUserClient(handler.userClientAddress)

	for _, id := range userIds {
		user, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: id})
		domainUser := mapNewUser(user.User)
		*users = append(*users, *domainUser)
	}
	return nil
}

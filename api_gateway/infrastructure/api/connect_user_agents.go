package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type ConnectUserAgentsHandler struct {
	userClientAddress string
}

type AgentsUser struct {
}

func NewConnectUserAgentsHandler(userClientAddress string) Handler {
	return &ConnectUserAgentsHandler{
		userClientAddress: userClientAddress,
	}
}

func (handler *ConnectUserAgentsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/connectWithDislinkt", handler.ConnectUser)
	if err != nil {
		panic(err)
	}
}

func (handler *ConnectUserAgentsHandler) ConnectUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var apiKey domain.ApiKey
	err := json.NewDecoder(r.Body).Decode(&apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userClient := services.NewUserClient(handler.userClientAddress)
	user, err := userClient.GetByUsername(context.TODO(), &user_proto.GetByUsernameRequest{Username: apiKey.Username})
	if err != nil {
		panic(err)
	}
	id := user.User.Id
	fmt.Println("ID:" + id)

	response, err := json.Marshal("Success")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

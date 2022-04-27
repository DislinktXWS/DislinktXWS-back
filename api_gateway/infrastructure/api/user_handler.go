package api

import (
	"context"
	"encoding/json"
	"log"
	"module/api_gateway/domain"
	"module/api_gateway/infrastructure/services"
	user_service "module/common/proto/user_service"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type UserHandler struct {
	userClientAddress string
}

func NewUserHandler(userClientAddress string) Handler {
	return &UserHandler{
		userClientAddress: userClientAddress,
	}
}

func (handler *UserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/users", handler.GetUsers)
	if err != nil {
		panic(err)
	}
}

func (handler *UserHandler) GetUsers(writer http.ResponseWriter, request *http.Request, pathParams map[string]string) {
	log.Println("Uhvatio handler")
	user := &domain.User{}

	err := handler.getRealUser(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

//prepraviti u getRealUserS
func (handler *UserHandler) getRealUser(user *domain.User) error {
	userClient := services.NewUserClient(handler.userClientAddress)
	realUser, err := userClient.Get(context.TODO(), &user_service.GetRequest{})
	if err != nil {
		return err
	}

	//setovanje polja
	user.Name = realUser.User.Name
	user.Surname = realUser.User.Surname

	return nil
}

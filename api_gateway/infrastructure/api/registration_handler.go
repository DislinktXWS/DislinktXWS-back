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

type RegistrationHandler struct {
	userClientAddress       string
	connectionClientAddress string
}

func NewRegistrationHandler(userClientAddress, connectionClientAddress string) Handler {
	return &RegistrationHandler{
		userClientAddress:       userClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *RegistrationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/users/user", handler.RegisterUser)
	if err != nil {
		panic(err)
	}
}

func (handler *RegistrationHandler) RegisterUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	var newUser domain.UserRegistration

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*newUserId*/
	newUserId, token, error := handler.addUser(newUser)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//samo dok se ne resi problem sa dockerom i neo4j je zakomentarisano, ogi <3

	e := handler.addUserNode(newUserId)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("*********************************************")
		fmt.Println(e)
		return
	}

	newUser.Id = newUserId
	newUser.VerificationToken = token
	response, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *RegistrationHandler) addUser(newUser domain.UserRegistration) (string, string, error) {

	userClient := services.NewUserClient(handler.userClientAddress)
	UserPb := mapToUserPb(&newUser)
	insertedUser, err := userClient.Insert(context.TODO(), &user_proto.InsertUserRequest{User: UserPb})

	return insertedUser.User.Id, insertedUser.User.VerificationToken, err
}
func (handler *RegistrationHandler) addUserNode(userId string) error {

	connectionClient := services.NewConnectionClient(handler.connectionClientAddress)
	_, err := connectionClient.InsertNewUser(context.TODO(), &connection_proto.InsertUserRequest{User: userId})

	return err
}

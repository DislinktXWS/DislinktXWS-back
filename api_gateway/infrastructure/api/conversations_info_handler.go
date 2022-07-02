package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	message_proto "github.com/dislinktxws-back/common/proto/message_service"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ConversationInfoHandler struct {
	userClientAddress    string
	messageClientAddress string
}

func NewConversationInfoHandler(userClientAddress, messageClientAddress string) Handler {
	return &ConversationInfoHandler{
		userClientAddress:    userClientAddress,
		messageClientAddress: messageClientAddress,
	}
}

func (handler *ConversationInfoHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/messages/{id}", handler.GetConversationsInfo)
	if err != nil {
		panic(err)
	}
}

func (handler *ConversationInfoHandler) GetConversationsInfo(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Print("ID je -> ")
	fmt.Print(id)

	result := []domain.ConversationInfo{}
	_ = handler.getConversationsInfoList(&result, id)

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *ConversationInfoHandler) getConversationsInfoList(users *[]domain.ConversationInfo, userId string) error {

	fmt.Println("USLI SMO U OVU JEBENU METODY")
	messagesClient := services.NewMessageClient(handler.messageClientAddress)
	conversationsPb, err := messagesClient.GetAllConversations(context.TODO(), &message_proto.GetAllConversationsRequest{Id: userId})

	if err != nil {
		fmt.Println("problem u gettovanju svih razgovora")
	}
	fmt.Println("DUzina")
	fmt.Println(len(conversationsPb.GetConversations()))

	for _, conversation := range conversationsPb.GetConversations() {

		fmt.Println("Pronadjen je conversation")
		fmt.Println(conversation.GetId())
		var otherUserId string

		if conversation.GetFirstParticipator() == userId {
			otherUserId = conversation.GetSecondParticipator()
		} else {
			otherUserId = conversation.GetFirstParticipator()
		}

		var userInfo *domain.UserBasicInfo
		handler.getUser(userInfo, otherUserId)

		conversationInfo := mapNewConversationInfo(conversation, userInfo)
		*users = append(*users, *conversationInfo)
	}
	return nil
}

func (handler *ConversationInfoHandler) getUser(user *domain.UserBasicInfo, userId string) error {

	userClient := services.NewUserClient(handler.userClientAddress)

	userById, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: userId})

	domainUser := mapNewUser(userById.User)
	*user = *domainUser

	return nil
}

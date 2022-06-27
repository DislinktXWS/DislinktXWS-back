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

type ConversationsSortedHandler struct {
	userClientAddress    string
	messageClientAddress string
}

func NewConversationsSortedHandler(userClientAddress, messageClientAddress string) Handler {
	return &ConversationsSortedHandler{
		userClientAddress:    userClientAddress,
		messageClientAddress: messageClientAddress,
	}
}

func (handler *ConversationsSortedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/conversations/{id}", handler.GetSortedConversations)
	if err != nil {
		panic(err)
	}
}

func (handler *ConversationsSortedHandler) GetSortedConversations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Print("ID je -> ")
	fmt.Print(id)

	result := []domain.ConversationInfo{}
	_ = handler.getConversationsInfo(&result, id)

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *ConversationsSortedHandler) getConversationsInfo(users *[]domain.ConversationInfo, userId string) error {

	conversations, _ := handler.getConversations(userId)

	for _, conversation := range conversations.Conversations {

		var otherUserId string

		if conversation.FirstParticipator == userId {
			otherUserId = conversation.SecondParticipator
		} else {
			otherUserId = conversation.FirstParticipator
		}

		var userInfo *domain.UserBasicInfo
		handler.getUser(userInfo, otherUserId)

		conversationInfo := mapNewConversationInfo(conversation, userInfo)
		*users = append(*users, *conversationInfo)
	}
	return nil
}

func (handler *ConversationsSortedHandler) getConversations(userId string) (*message_proto.GetAllConversationsResponse, error) {

	messagesClient := services.NewMessageClient(handler.messageClientAddress)
	conversations, err := messagesClient.GetAllConversations(context.TODO(), &message_proto.GetAllConversationsRequest{Id: userId})
	return conversations, err
}

func (handler *ConversationsSortedHandler) getUser(user *domain.UserBasicInfo, userId string) error {

	userClient := services.NewUserClient(handler.userClientAddress)

	userById, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: userId})

	domainUser := mapNewUser(userById.User)
	*user = *domainUser

	return nil
}

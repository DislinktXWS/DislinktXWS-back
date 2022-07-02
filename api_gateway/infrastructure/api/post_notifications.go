package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	notifications_proto "github.com/dislinktxws-back/common/proto/notifications_service"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type PostNotificationsHandler struct {
	notificationClientAddress string
	connectionsClientAddress  string
	userClientAddress         string
}

func NewPostNotificationsHandler(notificationClientAddress string, connectionsClientAddress string, userClientAddress string) Handler {
	return &PostNotificationsHandler{
		notificationClientAddress: notificationClientAddress,
		connectionsClientAddress:  connectionsClientAddress,
		userClientAddress:         userClientAddress,
	}
}

func (handler *PostNotificationsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/notifications/postNotifications/{id}", handler.PostNotifications)
	if err != nil {
		panic(err)
	}
}

func (handler *PostNotificationsHandler) PostNotifications(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	fmt.Println("POGODJEN POST NOTIF HANDLER")
	id := pathParams["id"]
	var newNotification domain.Notification
	err := json.NewDecoder(r.Body).Decode(&newNotification)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sender := domain.UserBasicInfo{}
	receiver := domain.UserBasicInfo{}
	sender = *handler.getUserInformation(id, &sender)
	receiver = *handler.getUserInformation(id, &receiver)
	handler.sendNotification(&sender, &receiver, newNotification)

}

func (handler *PostNotificationsHandler) getUserInformation(id string, userInfo *domain.UserBasicInfo) *domain.UserBasicInfo {
	userClient := services.NewUserClient(handler.userClientAddress)
	user, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: id})
	userInfo = mapNewUser(user.User)
	return userInfo
}

func (handler *PostNotificationsHandler) sendNotification(sender *domain.UserBasicInfo, receiver *domain.UserBasicInfo, notification domain.Notification) error {
	notificationsClient := services.NewNotificationClient(handler.notificationClientAddress)
	notificationsClient.Insert(context.TODO(), &notifications_proto.InsertRequest{Notification: mapNotificationPb(&notification, mapUserPb(sender), mapUserPb(receiver))})
	return nil
}

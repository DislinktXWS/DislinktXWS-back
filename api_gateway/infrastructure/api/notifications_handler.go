package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	connection_proto "github.com/dislinktxws-back/common/proto/connection_service"
	notifications_proto "github.com/dislinktxws-back/common/proto/notifications_service"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type NotificationHandler struct {
	notificationClientAddress string
	connectionsClientAddress  string
	userClientAddress         string
}

func NewNotificationHandler(notificationClientAddress string, connectionsClientAddress string, userClientAddress string) Handler {
	return &NotificationHandler{
		notificationClientAddress: notificationClientAddress,
		connectionsClientAddress:  connectionsClientAddress,
		userClientAddress:         userClientAddress,
	}
}

func (handler *NotificationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/notifications/notifyConnections/{id}", handler.NotifyConnections)
	if err != nil {
		panic(err)
	}
}

func (handler *NotificationHandler) NotifyConnections(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var newNotification domain.Notification
	id := pathParams["id"]
	fmt.Println("ID SENDERA")
	fmt.Println(id)

	err := json.NewDecoder(r.Body).Decode(&newNotification)
	fmt.Println("OBAVESTENJE PROSLEDJENO SA FRONTA")
	fmt.Println(newNotification)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userConnections, err := handler.getUserConnections(id)
	fmt.Println("KONEKCIJE USERA")
	fmt.Println(userConnections)

	users := []domain.UserBasicInfo{}
	sender := domain.UserBasicInfo{}

	sender = *handler.getSenderInformation(id, &sender)
	fmt.Println("SENDER INFORMATION:")
	fmt.Println(sender)

	handler.getUsersToNotify(userConnections, &users)

	handler.sendNotifications(&sender, &users, newNotification)

	fmt.Println(users)
}

func (handler *NotificationHandler) getUserConnections(userId string) ([]string, error) {
	connectionsClient := services.NewConnectionClient(handler.connectionsClientAddress)
	connections, err := connectionsClient.GetAll(context.TODO(), &connection_proto.GetAllConnectionsRequest{Id: userId})
	return connections.Ids, err
}

func (handler *NotificationHandler) getSenderInformation(id string, sender *domain.UserBasicInfo) *domain.UserBasicInfo {
	userClient := services.NewUserClient(handler.userClientAddress)
	user, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: id})
	fmt.Println(user.User)
	sender = mapNewUser(user.User)
	fmt.Println("SENDEr nakon dobavlajanja")
	fmt.Println(sender)
	return sender
}

func (handler *NotificationHandler) getUsersToNotify(userIds []string, users *[]domain.UserBasicInfo) *[]domain.UserBasicInfo {
	userClient := services.NewUserClient(handler.userClientAddress)

	for _, id := range userIds {
		user, _ := userClient.Get(context.TODO(), &user_proto.GetRequest{Id: id})
		domainUser := mapNewUser(user.User)
		fmt.Println("NABAVI PODESAVANJA OBAVESTENJA OD USERA:" + id)
		userNotificationSettings, _ := userClient.GetNotificationsSettings(context.TODO(), &user_proto.GetNotificationsSettingsRequest{Id: id})
		fmt.Println(userNotificationSettings)
		if userNotificationSettings.ConnectionsNotifications {
			*users = append(*users, *domainUser)
		}
	}
	return users
}

func (handler *NotificationHandler) sendNotifications(sender *domain.UserBasicInfo, users *[]domain.UserBasicInfo, notification domain.Notification) error {
	notificationsClient := services.NewNotificationClient(handler.notificationClientAddress)
	for _, user := range *users {
		notificationsClient.Insert(context.TODO(),
			&notifications_proto.InsertRequest{
				Notification: mapNotificationPb(&notification, mapUserPb(sender), mapUserPb(&user))})
	}

	return nil
}

func mapNotificationPb(notification *domain.Notification, userFrom *notifications_proto.User, userTo *notifications_proto.User) *notifications_proto.Notification {

	notificationPb := &notifications_proto.Notification{
		Id:      notification.Id,
		From:    userFrom,
		To:      userTo,
		Date:    notification.Date,
		Content: notification.Content,
	}
	return notificationPb
}

func mapUserPb(user *domain.UserBasicInfo) *notifications_proto.User {
	userPb := &notifications_proto.User{
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
	}
	return userPb
}

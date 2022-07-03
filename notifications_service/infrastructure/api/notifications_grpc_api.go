package api

import (
	"context"
	"fmt"
	pb "github.com/dislinktxws-back/common/proto/notifications_service"
	"github.com/dislinktxws-back/notification_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationsServiceServer
	service *application.NotificationsService
}

func NewNotificationsHandler(service *application.NotificationsService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (handler *NotificationHandler) Insert(ctx context.Context, request *pb.InsertRequest) (*pb.InsertResponse, error) {
	notification := mapNewNotification(request.Notification)
	fmt.Println("USLO U METODU INSERT")
	fmt.Println(notification)
	newNotification, err := handler.service.Insert(notification)
	fmt.Println("IZASLO IZ METODE INSERT")
	if err != nil {
		return nil, err
	}
	return &pb.InsertResponse{Notification: mapNewNotificationPb(newNotification)}, nil
}

func (handler *NotificationHandler) GetUserNotifications(ctx context.Context, request *pb.GetUserNotificationsRequest) (*pb.GetUserNotificationsResponse, error) {
	username := request.Username
	notifications, _ := handler.service.GetUserNotifications(username)
	response := &pb.GetUserNotificationsResponse{
		NotificationView: []*pb.NotificationView{},
	}
	for _, notification := range notifications {
		current := mapNotification(notification)
		response.NotificationView = append(response.NotificationView, current)
	}

	return response, nil
}

func (handler *NotificationHandler) ReviewNotification(ctx context.Context, request *pb.ReviewNotificationRequest) (*pb.ReviewNotificationResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	err := handler.service.ReviewNotification(id)
	if err != nil {
		fmt.Println("ERROR with UPDATE notification!")
		return nil, err
	}
	return &pb.ReviewNotificationResponse{}, nil
}

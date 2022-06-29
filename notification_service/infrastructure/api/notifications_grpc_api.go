package api

import (
	pb "github.com/dislinktxws-back/common/proto/notifications_service"
	"github.com/dislinktxws-back/notification_service/application"
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

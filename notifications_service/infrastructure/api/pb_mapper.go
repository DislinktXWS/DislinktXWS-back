package api

import (
	pb "github.com/dislinktxws-back/common/proto/notifications_service"
	"github.com/dislinktxws-back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapNewNotification(notficationPb *pb.Notification) *domain.Notification {
	userFrom := domain.User{
		Username: notficationPb.From.Username,
		Name:     notficationPb.From.Name,
		Surname:  notficationPb.From.Surname,
	}
	userTo := domain.User{
		Username: notficationPb.To.Username,
		Name:     notficationPb.To.Name,
		Surname:  notficationPb.To.Surname,
	}
	notification := &domain.Notification{
		Id:      primitive.ObjectID{},
		From:    userFrom,
		To:      userTo,
		Date:    notficationPb.Date,
		Content: notficationPb.Content,
	}
	return notification
}

func mapNewNotificationPb(notification *domain.Notification) *pb.Notification {
	userFromPb := &pb.User{
		Username: notification.From.Username,
		Name:     notification.From.Name,
		Surname:  notification.From.Surname,
	}
	userToPb := &pb.User{
		Username: notification.To.Username,
		Name:     notification.To.Name,
		Surname:  notification.To.Surname,
	}
	notificationPb := &pb.Notification{
		Id:      notification.Id.Hex(),
		From:    userFromPb,
		To:      userToPb,
		Date:    notification.Date,
		Content: notification.Content,
	}

	return notificationPb
}

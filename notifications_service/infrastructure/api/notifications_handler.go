package api

import "github.com/dislinktxws-back/notification_service/application"

type CreateNotificationsCommandHandler struct {
	notificationsService *application.NotificationsService
}

func NewCreateUserCommandHandler(notificationService *application.NotificationsService) (*CreateNotificationsCommandHandler, error) {
	o := &CreateNotificationsCommandHandler{
		notificationsService: notificationService,
	}
	return o, nil
}

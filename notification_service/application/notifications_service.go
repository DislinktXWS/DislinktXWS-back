package application

import "github.com/dislinktxws-back/notification_service/domain"

type NotificationsService struct {
	store domain.NotificationsStore
}

func NewNotificationsService(store domain.NotificationsStore) *NotificationsService {
	return &NotificationsService{
		store: store,
	}
}

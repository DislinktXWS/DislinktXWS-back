package application

import (
	"github.com/dislinktxws-back/notification_service/domain"
)

type NotificationsService struct {
	store domain.NotificationsStore
}

func (service *NotificationsService) Insert(notification *domain.Notification) (*domain.Notification, error) {
	return service.store.Insert(notification)
}

func (service *NotificationsService) GetUserNotifications(username string) ([]*domain.Notification, error) {
	return service.store.GetUserNotifications(username)
}

func NewNotificationsService(store domain.NotificationsStore) *NotificationsService {
	return &NotificationsService{
		store: store,
	}
}

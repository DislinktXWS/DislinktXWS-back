package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NotificationsStore interface {
	Get(id primitive.ObjectID) error
	Insert(notification *Notification) (*Notification, error)
	GetUserNotifications(username string) ([]*Notification, error)
	ReviewNotification(id primitive.ObjectID) error
}

package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type NotificationsStore interface {
	Get(id primitive.ObjectID) error
}

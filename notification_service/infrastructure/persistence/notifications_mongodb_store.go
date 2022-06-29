package persistence

import (
	"github.com/dislinktxws-back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "notificationsDB"
	COLLECTION = "notifications"
)

type NotificationsMongoDBStore struct {
	notifications *mongo.Collection
}

func (n NotificationsMongoDBStore) Get(id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func NewNotificationsMongoDBStore(client *mongo.Client) domain.NotificationsStore {
	notificaitions := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationsMongoDBStore{
		notifications: notificaitions,
	}
}

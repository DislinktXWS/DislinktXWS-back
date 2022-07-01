package persistence

import (
	"context"
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

func (store *NotificationsMongoDBStore) Insert(notification *domain.Notification) (*domain.Notification, error) {
	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return &domain.Notification{}, err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return notification, nil
}

func (store *NotificationsMongoDBStore) Get(id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func NewNotificationsMongoDBStore(client *mongo.Client) domain.NotificationsStore {
	notificaitions := client.Database(DATABASE).Collection(COLLECTION)
	return &NotificationsMongoDBStore{
		notifications: notificaitions,
	}
}

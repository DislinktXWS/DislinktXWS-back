package persistence

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/notification_service/domain"
	"go.mongodb.org/mongo-driver/bson"
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

func (store *NotificationsMongoDBStore) ReviewNotification(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	notification, err := store.filterOne(filter)
	if err != nil {
		fmt.Println("Notification with id: %s doesn't exist!", id)
		return err
	}
	_, err1 := store.notifications.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"isReviewed", !notification.IsReviewed}}},
		},
	)
	return err1

}

func (store *NotificationsMongoDBStore) filterOne(filter interface{}) (Notification *domain.Notification, err error) {
	result := store.notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&Notification)
	return
}

func (store *NotificationsMongoDBStore) GetUserNotifications(username string) ([]*domain.Notification, error) {
	filter := bson.D{{}}
	allNotifications, _ := store.filter(filter)
	fmt.Println(allNotifications)
	notifications := []*domain.Notification{}
	for _, notification := range allNotifications {
		if notification.To.Username == username {
			notifications = append(notifications, notification)
		}
	}
	return notifications, nil
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

func (store *NotificationsMongoDBStore) filter(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (posts []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var Notification domain.Notification
		err = cursor.Decode(&Notification)
		if err != nil {
			return
		}
		posts = append(posts, &Notification)
	}
	err = cursor.Err()
	return
}

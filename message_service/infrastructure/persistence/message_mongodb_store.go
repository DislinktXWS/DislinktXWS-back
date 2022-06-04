package persistence

import (
	"github.com/dislinktxws-back/message_service/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "messageDB"
	COLLECTION = "messages"
)

type MessageMongoDBStore struct {
	messages *mongo.Collection
}

func NewMessageMongoDBStore(client *mongo.Client) domain.MessageStore {
	messages := client.Database(DATABASE).Collection(COLLECTION)
	return &MessageMongoDBStore{
		messages: messages,
	}
}

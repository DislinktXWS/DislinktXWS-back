package persistence

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/message_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "messDB"
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

func (store *MessageMongoDBStore) CreateConversation(participants *domain.Participants) error {

	filter := bson.M{
		"$or": []bson.M{
			{"$and": []bson.M{
				{"first_participator": participants.Receiver},
				{"second_participator": participants.Sender},
			}},
			{"$and": []bson.M{
				{"first_participator": participants.Sender},
				{"second_participator": participants.Receiver},
			}},
		},
	}

	_, e := store.filterOne(filter)

	if e != nil {

		Conversation := new(domain.Conversation)
		Conversation.Messages = make([]domain.Message, 0)
		Conversation.FirstParticipator = participants.Sender
		Conversation.SecondParticipator = participants.Receiver

		_, err := store.messages.InsertOne(context.TODO(), Conversation)
		if err != nil {
			return err
		}

		fmt.Println("DOSAO JE DO KRAJA STORE METODE")
	}
	return nil
}

func (store *MessageMongoDBStore) GetAllConversations(userId string) ([]*domain.Conversation, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"first_participator": userId},
			{"second_participator": userId},
		}}
	return store.filter(filter)
}

func (store *MessageMongoDBStore) AddMessage(message *domain.Message) error {

	participants := new(domain.Participants)
	participants.Sender = message.Sender
	participants.Receiver = message.Receiver
	messageHistory, err := store.GetConversation(participants)

	message.IsRead = false

	if messageHistory == nil {
		fmt.Println("Ne postoji conversation za ovu gospodu")
	} else {
		messages := append(messageHistory.Messages, *message)

		_, err := store.messages.UpdateOne(context.TODO(), bson.M{"_id": messageHistory.Id}, bson.D{
			{"$set", bson.D{{"messages", messages}}},
		})
		if err != nil {
			return err
		}
	}
	return err
}

func (store *MessageMongoDBStore) GetConversation(participants *domain.Participants) (*domain.Conversation, error) {

	filter := bson.M{
		"$or": []bson.M{
			{"$and": []bson.M{
				{"first_participator": participants.Receiver},
				{"second_participator": participants.Sender},
			}},
			{"$and": []bson.M{
				{"first_participator": participants.Sender},
				{"second_participator": participants.Receiver},
			}},
		},
	}

	conversation, err := store.filterOne(filter)

	for index, message := range conversation.Messages {
		if message.Receiver == participants.Sender {
			conversation.Messages[index].IsRead = true
		}
	}

	store.messages.UpdateOne(
		context.TODO(),
		bson.M{"_id": conversation.Id},
		bson.D{
			{"$set", bson.D{
				{"messages", conversation.Messages},
			}},
		},
	)
	return conversation, err
}

func (store *MessageMongoDBStore) filter(filter interface{}) ([]*domain.Conversation, error) {
	cursor, err := store.messages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *MessageMongoDBStore) filterOne(filter interface{}) (Conversation *domain.Conversation, err error) {
	result := store.messages.FindOne(context.TODO(), filter)
	err = result.Decode(&Conversation)
	return
}

func decode(cursor *mongo.Cursor) (orders []*domain.Conversation, err error) {
	for cursor.Next(context.TODO()) {
		var Conversation domain.Conversation
		err = cursor.Decode(&Conversation)
		if err != nil {
			return
		}
		orders = append(orders, &Conversation)
	}
	err = cursor.Err()
	return
}

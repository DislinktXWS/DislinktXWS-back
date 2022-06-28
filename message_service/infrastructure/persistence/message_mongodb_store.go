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
	
	existingConversation, _ := store.GetConversation(*participants)
	if existingConversation != nil {
		return nil
	} else {

		Conversation := new(domain.Conversation)
		Conversation.Messages = make([]domain.Message, 0)
		Conversation.FirstParticipator = participants.Sender
		Conversation.SecondParticipator = participants.Receiver

		_, err := store.messages.InsertOne(context.TODO(), Conversation)
		if err != nil {
			return err
		}

		fmt.Println("DOSAO JE DO KRAJA STORE METODE")
		return nil
	}
}

func (store *MessageMongoDBStore) GetAllConversations(userId string) ([]*domain.Conversation, error) {

	var conversations []string
	var conversationsObject []*domain.Conversation

	filterFirst := bson.M{"first_participator": userId}
	resultFirst, errFirst := store.filter(filterFirst)

	if errFirst == nil {
		for index, element := range resultFirst {
			fmt.Println("At index", index, "value is", element.SecondParticipator)

			if !contains(conversations, element.SecondParticipator) {
				conversations = append(conversations, element.SecondParticipator)
				conversationsObject = append(conversationsObject, element)
			}
		}
	}

	filterSecond := bson.M{"second_participator": userId}
	resultSecond, errSecond := store.filter(filterSecond)

	if errSecond == nil {
		for index, element := range resultSecond {
			fmt.Println("At index", index, "value is", element.SecondParticipator)

			if !contains(conversations, element.SecondParticipator) {
				conversations = append(conversations, element.SecondParticipator)
				conversationsObject = append(conversationsObject, element)
			}
		}
	}
	return conversationsObject, nil
}

func (store *MessageMongoDBStore) AddMessage(message *domain.Message, participants domain.Participants) error {

	conversation, _ := store.GetConversation(participants)

	messages := conversation.Messages
	messages = append(messages, *message)

	_, err := store.messages.UpdateOne(
		context.TODO(),
		bson.M{"_id": conversation.Id},
		bson.D{
			{"$set", bson.D{{"messages", messages}}},
		},
	)
	return err
}

func (store *MessageMongoDBStore) GetConversation(participants domain.Participants) (Conversation *domain.Conversation, err error) {

	filter := bson.M{
		"first_participator":  participants.Sender,
		"second_participator": participants.Receiver,
	}

	e := store.messages.FindOne(context.TODO(), filter).Decode(&Conversation)

	if e != nil {

		filter := bson.M{
			"first_participator":  participants.Receiver,
			"second_participator": participants.Sender,
		}
		eSecond := store.messages.FindOne(context.TODO(), filter).Decode(&Conversation)

		if eSecond != nil {
			return nil, nil
		}
	}
	fmt.Println("USTVARI JE PROSAO STORE")
	return
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

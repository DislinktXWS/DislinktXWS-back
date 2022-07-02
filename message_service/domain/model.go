package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Conversation struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	FirstParticipator  string             `bson:"first_participator"`
	SecondParticipator string             `bson:"second_participator"`
	Messages           []Message          `bson:"messages"`
	LastMessageDate    time.Time          `bson:"last_message_date"`
}

type Message struct {
	Sender   string    `bson:"sender"`
	Receiver string    `bson:"receiver"`
	Content  string    `bson:"content"`
	Date     time.Time `bson:"date"`
	IsRead   bool      `bson:"is_read"`
}

type Participants struct {
	Sender   string
	Receiver string
}

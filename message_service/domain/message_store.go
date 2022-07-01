package domain

type MessageStore interface {
	CreateConversation(participants *Participants) error
	GetConversation(participants *Participants) (*Conversation, error)
	GetAllConversations(userId string) ([]*Conversation, error)
	AddMessage(message *Message) error
}

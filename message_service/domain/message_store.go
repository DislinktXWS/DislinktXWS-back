package domain

type MessageStore interface {
	CreateConversation(participants Participants) (error, *Conversation)
	GetConversation(participants Participants) (*Conversation, error)
	GetAllConversations(userId string) ([]*Conversation, error)
	AddMessage(message *Message, participants Participants) error
}

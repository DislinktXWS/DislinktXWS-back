package domain

type MessageStore interface {
	CreateConversation(participants Participants) (error, *Conversation)
	GetConversation(participants Participants) (*Conversation, error)
	GetAllConversations(userId string) ([]string, error)
	AddMessage(skill *Message, participants Participants) error
}

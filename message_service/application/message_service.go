package application

import (
	"github.com/dislinktxws-back/message_service/domain"
)

type MessageService struct {
	store domain.MessageStore
}

func NewMessageService(store domain.MessageStore) *MessageService {
	return &MessageService{
		store: store,
	}
}

func (service *MessageService) CreateConversation(participants *domain.Participants) error {
	return service.store.CreateConversation(participants)
}
func (service *MessageService) GetConversation(participants *domain.Participants) (*domain.Conversation, error) {
	return service.store.GetConversation(participants)
}
func (service *MessageService) GetAllConversations(userId string) ([]*domain.Conversation, error) {
	return service.store.GetAllConversations(userId)
}
func (service *MessageService) AddMessage(message *domain.Message) error {
	return service.store.AddMessage(message)
}

package api

import (
	"github.com/dislinktxws-back/message_service/application"
)

type CreateMessageCommandHandler struct {
	messageService *application.MessageService
}

func NewCreateUserCommandHandler(messageService *application.MessageService) (*CreateMessageCommandHandler, error) {
	o := &CreateMessageCommandHandler{
		messageService: messageService,
	}
	return o, nil
}

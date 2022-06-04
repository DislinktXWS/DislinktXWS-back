package api

import (
	pb "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/dislinktxws-back/message_service/application"
)

type MessageHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.MessageService
}

func NewMessageHandler(service *application.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

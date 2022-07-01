package api

import (
	"context"
	"fmt"
	pb "github.com/dislinktxws-back/common/proto/message_service"
	"github.com/dislinktxws-back/message_service/application"
	"github.com/dislinktxws-back/message_service/domain"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer
	service *application.MessageService
}

func NewMessageHandler(service *application.MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (handler *MessageHandler) CreateConversation(ctx context.Context, request *pb.CreateConversationRequest) (*pb.EmptyResponse, error) {

	participants := mapNewConversation(request.Participants)
	fmt.Println("OVO JE REZ MAPIRANJA")
	fmt.Println(participants.Sender)
	fmt.Println(participants.Receiver)
	err := handler.service.CreateConversation(participants)

	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

func (handler *MessageHandler) GetConversation(ctx context.Context, request *pb.GetConversationRequest) (*pb.GetConversationResponse, error) {

	fmt.Print("DESILA SE METODA GET CONVERSATION")

	participants := new(domain.Participants)
	participants.Sender = request.GetId1()
	participants.Receiver = request.GetId2()

	conversation, err := handler.service.GetConversation(participants)
	if conversation == nil {
		fmt.Println("CONVERSATION JE NULL")
		return &pb.GetConversationResponse{}, nil
	}

	conversationPb := mapConversation(conversation)

	if err != nil {
		return nil, err
	}
	return conversationPb, nil
}

func (handler *MessageHandler) GetAllConversations(ctx context.Context, request *pb.GetAllConversationsRequest) (*pb.GetAllConversationsResponse, error) {

	conversations, err := handler.service.GetAllConversations(request.Id)

	response := &pb.GetAllConversationsResponse{}

	for _, conversation := range conversations {
		response.Conversations = append(response.Conversations, mapConversation(conversation))
	}
	return response, err
}

func (handler *MessageHandler) AddMessage(ctx context.Context, request *pb.AddMessageRequest) (*pb.EmptyResponse, error) {

	fmt.Print("DESILA SE METODA")

	message := mapNewMessage(request.Message)
	err := handler.service.AddMessage(message)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

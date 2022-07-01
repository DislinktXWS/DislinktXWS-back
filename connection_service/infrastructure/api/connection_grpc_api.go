package api

import (
	"context"
	"fmt"
	connections "github.com/dislinktxws-back/common/proto/connection_service"
	pb "github.com/dislinktxws-back/common/proto/connection_service"
	events "github.com/dislinktxws-back/common/saga/insert_user"
	saga "github.com/dislinktxws-back/common/saga/messaging"
	"github.com/dislinktxws-back/connection_service/application"
	"log"
	"os"
)

type ConnectionHandler struct {
	pb.UnimplementedConnectionsServiceServer
	service           *application.ConnectionsService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func NewConnectionHandler(service *application.ConnectionsService, publisher saga.Publisher, subscriber saga.Subscriber) *ConnectionHandler {
	o := &ConnectionHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	o.commandSubscriber.Subscribe(o.handle)
	return o
}

func init() {
	infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(infoFile, "INFO: ", log.LstdFlags|log.Lshortfile)

	errFile, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	ErrorLogger = log.New(errFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func (handler *ConnectionHandler) handle(command *events.InsertUserCommand) {
	reply := events.InsertUserReply{User: command.User}
	fmt.Println("CONNECTION HANDLER")
	fmt.Println(command.Type)

	switch command.Type {
	case events.InsertUserNode:
		fmt.Println("INSERT NODE")
		err := handler.service.InsertNewUser(command.User.Id)
		if err != nil {
			reply.Type = events.UserNodeNotInserted
			break
		}
		reply.Type = events.UserNodeInserted
		//reply.Type = events.UserNodeNotInserted
	default:
		reply.Type = events.UnknownReply
	}

	fmt.Println(reply.Type)
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}

func (handler *ConnectionHandler) InsertUserConnection(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	fmt.Print("DESILA SE METODA")
	connection := mapNewUserConnection(request.Connection)
	err := handler.service.InsertUserConnection(connection)
	if err != nil {
		ErrorLogger.Println("Action: 15, Message: Cannot create connection!")
		return nil, err
	}
	InfoLogger.Println("Action: 16, Message: New connection between " + connection.Connected + " and " + connection.Connecting)
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) DeleteUserConnection(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.DeleteUserConnection(connection)
	if err != nil {
		ErrorLogger.Println("Action: 17, Cannot delete connection!")
		return nil, err
	}
	InfoLogger.Println("Message: 18, Deleted connection between " + connection.Connected + " and " + connection.Connecting)
	return &pb.InsertUserConnectionResponse{}, nil
}
func (handler *ConnectionHandler) InsertNewUser(ctx context.Context, request *pb.InsertUserRequest) (*pb.InsertUserResponse, error) {
	user := request.User
	err := handler.service.InsertNewUser(user)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Cannot create user!")
		return nil, err
	}
	InfoLogger.Println("Action: 3, Message: User " + user + " created.")
	return &pb.InsertUserResponse{}, nil
}

func (handler *ConnectionHandler) GetAll(ctx context.Context, request *pb.GetAllConnectionsRequest) (*pb.GetAllConnectionsResponse, error) {

	Connections := handler.service.GetAll(request.Id)

	response := &pb.GetAllConnectionsResponse{}

	for _, connection := range Connections {
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}
func (handler *ConnectionHandler) GetBlockedUsers(ctx context.Context, request *pb.GetAllConnectionsRequest) (*pb.GetAllConnectionsResponse, error) {

	Connections := handler.service.GetBlockedUsers(request.Id)

	response := &pb.GetAllConnectionsResponse{}

	for _, connection := range Connections {
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}
func (handler *ConnectionHandler) GetConnectionRequests(ctx context.Context, request *pb.GetAllConnectionsRequest) (*pb.GetAllConnectionsResponse, error) {

	Connections := handler.service.GetAllConnectionRequests(request.Id)

	response := &pb.GetAllConnectionsResponse{}

	for _, connection := range Connections {
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}

func (handler *ConnectionHandler) GetUserRecommendations(ctx context.Context, request *pb.GetAllConnectionsRequest) (*pb.GetAllConnectionsResponse, error) {

	Connections := handler.service.GetUserRecommendations(request.Id)

	response := &pb.GetAllConnectionsResponse{}

	for _, connection := range Connections {
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}
func (handler *ConnectionHandler) GetConnectionStatus(ctx context.Context, request *pb.ConnectionStatusRequest) (*pb.ConnectionStatusResponse, error) {

	var enums int32
	status := handler.service.GetConnectionStatus(request.Id1, request.Id2)
	if status == "connected" {
		enums = 0
	}
	if status == "connectionRequestedByYou" {
		enums = 1
	}
	if status == "connectionRequestedByUser" {
		enums = 2
	}
	if status == "blockedYou" {
		enums = 3
	}
	if status == "blockedByYou" {
		enums = 4
	}
	if status == "none" {
		enums = 5
	}
	response := &pb.ConnectionStatusResponse{Status: connections.ConnectionStatusEnum(enums)}
	return response, nil
}
func (handler *ConnectionHandler) InsertConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.InsertConnectionRequest(connection)
	if err != nil {
		ErrorLogger.Println("Action: 15, Message: Cannot create connection request!")
		return nil, err
	}
	InfoLogger.Println("Action: 19, Message: New connection request created between " + connection.Connected + " and " + connection.Connecting)
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) CancelConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.CancelConnectionRequest(connection)
	if err != nil {
		ErrorLogger.Println("Action: 16, Message: Cannot cancel connection request!")
		return nil, err
	}
	InfoLogger.Println("Action: 20, Message: Canceled connection request between " + connection.Connected + " and " + connection.Connecting)
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) BlockUser(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {
	connection := mapNewUserConnection(request.Connection)
	err := handler.service.BlockUser(connection)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Cannot block user!")
		return nil, err
	}
	InfoLogger.Println("Action: 21, Message: User " + connection.Connecting + " blocked " + connection.Connected)
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) UnblockUser(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.UnblockUser(connection)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Cannot unblock user!")
		return nil, err
	}
	InfoLogger.Println("Action: 22, Message: User " + connection.Connecting + " unblocked " + connection.Connected)

	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) AcceptConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.AcceptUserConnection(connection)
	if err != nil {
		ErrorLogger.Println("Action: 15, Message: Cannot accept connection request!")
	}
	InfoLogger.Println("Action: 23, Message: User " + connection.Connecting + " accepted connection request for " + connection.Connected)
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) DeclineConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.DeclineUserConnection(connection)
	if err != nil {
		ErrorLogger.Println("Action: 15, Message: Cannot decline connection request!")
		return nil, err
	}
	InfoLogger.Println("Action: 24, Message: User " + connection.Connecting + " declined connection request for " + connection.Connected)
	return &pb.InsertUserConnectionResponse{}, nil
}

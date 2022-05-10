package api

import (
	"context"
	"fmt"
	connections "github.com/dislinktxws-back/common/proto/connection_service"
	pb "github.com/dislinktxws-back/common/proto/connection_service"
	"github.com/dislinktxws-back/connection_service/application"
)

type ConnectionHandler struct {
	pb.UnimplementedConnectionsServiceServer
	service *application.ConnectionsService
}

func NewConnectionHandler(service *application.ConnectionsService) *ConnectionHandler {
	return &ConnectionHandler{
		service: service,
	}
}
func (handler *ConnectionHandler) InsertUserConnection(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	fmt.Print("DESILA SE METODA")
	connection := mapNewUserConnection(request.Connection)
	err := handler.service.InsertUserConnection(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) DeleteUserConnection(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.DeleteUserConnection(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}
func (handler *ConnectionHandler) InsertNewUser(ctx context.Context, request *pb.InsertUserRequest) (*pb.InsertUserResponse, error) {
	user := request.User
	err := handler.service.InsertNewUser(user)
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) CancelConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.CancelConnectionRequest(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) BlockUser(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.BlockUser(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) UnblockUser(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.UnblockUser(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) AcceptConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.AcceptUserConnection(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

func (handler *ConnectionHandler) DeclineConnectionRequest(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserConnectionResponse, error) {

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.DeclineUserConnection(connection)
	if err != nil {
		return nil, err
	}
	return &pb.InsertUserConnectionResponse{}, nil
}

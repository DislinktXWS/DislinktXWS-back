package api

import (
	"context"
	pb "module/common/proto/connection_service"
	"module/connection_service/application"
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

	connection := mapNewUserConnection(request.Connection)
	err := handler.service.InsertUserConnection(connection)
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

	BlockUsers := handler.service.GetBlockedUsers(request.Id)

	response := &pb.GetAllConnectionsResponse{}

	for _, connection := range BlockUsers {
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}

func (handler *ConnectionHandler) GetConnectionRequests(ctx context.Context, request *pb.GetAllConnectionsRequest) (*pb.GetAllConnectionsResponse, error) {

	ConnectionRequests := handler.service.GetAllConnectionRequests(request.Id)

	response := &pb.GetAllConnectionsResponse{}
	for _, connection := range ConnectionRequests {
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}

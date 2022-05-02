package api

import (
	"context"
	"fmt"
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
		fmt.Print("         KOnekcija jedna         ")
		fmt.Print(connection)
		response.Ids = append(response.Ids, connection)
	}
	return response, nil
}

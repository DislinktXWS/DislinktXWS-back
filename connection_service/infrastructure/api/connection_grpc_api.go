package api

import (
	"context"
	pb "module/common/proto/connection_service"
	"module/connection_service/application"
)

type ConnectionHandler struct {
	//pb.UnimplementedConnectionServiceServer
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

func (handler *ConnectionHandler) InsertNewUser(ctx context.Context, request *pb.InsertUserConnectionRequest) (*pb.InsertUserResponse, error) {
	user := request.String()
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

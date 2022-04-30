package api

import (
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

////potrebno izgenerisati proto da bi se napisali handleri uopste

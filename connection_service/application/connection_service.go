package application

import (
	"module/connection_service/domain"
)

type ConnectionsService struct {
	graph domain.ConnectionsGraph
}

func NewConnectionsService(graph domain.ConnectionsGraph) *ConnectionsService {
	return &ConnectionsService{
		graph: graph,
	}
}

/*
func (service *ConnectionsService) Get(connection *domain.UserConnection) bool {
	return service.graph.Get(connection)
}*/

func (service *ConnectionsService) GetAll(userId string) []string {
	return service.graph.GetAll(userId)
}
func (service *ConnectionsService) InsertNewUser(user string) error {
	return service.graph.InsertNewUser(user)
}

func (service *ConnectionsService) InsertUserConnection(connection *domain.UserConnection) error {
	return service.graph.InsertUserConnection(connection)
}

func (service *ConnectionsService) Delete(connection *domain.UserConnection) error {
	return service.graph.DeleteUserConnection(connection)
}

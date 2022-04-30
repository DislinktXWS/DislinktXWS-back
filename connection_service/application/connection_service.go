package application

import "module/connection_service/domain"

type ConnectionsService struct {
	graph domain.ConnectionsGraph
}

func NewConnectionsService(graph domain.ConnectionsGraph) *ConnectionsService {
	return &ConnectionsService{
		graph: graph,
	}
}

func (service *ConnectionsService) Get(connection *domain.UserConnection) bool {
	return service.graph.Get(connection)
}

func (service *ConnectionsService) GetAll(username string) []*domain.User {
	return service.graph.GetAll(username)
}
func (service *ConnectionsService) InsertNewUser(user *domain.User) error {
	return service.graph.InsertNewUser(user)
}

func (service *ConnectionsService) InsertUserConnection(connection *domain.UserConnection) error {
	return service.graph.InsertUserConnection(connection)
}

func (service *ConnectionsService) Delete(connection *domain.UserConnection) error {
	return service.graph.DeleteUserConnection(connection)
}

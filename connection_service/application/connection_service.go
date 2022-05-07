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

func (service *ConnectionsService) GetAll(userId string) []string {
	return service.graph.GetAll(userId)
}
func (service *ConnectionsService) GetAllConnectionRequests(userId string) []string {
	return service.graph.GetConnectionRequests(userId)
}
func (service *ConnectionsService) GetConnectionStatus(user1 string, user2 string) string {
	return service.graph.GetConnectionStatus(user1, user2)
}
func (service *ConnectionsService) GetBlockedUsers(userId string) []string {
	return service.graph.GetBlockedUsers(userId)
}

func (service *ConnectionsService) InsertNewUser(user string) error {
	return service.graph.InsertNewUser(user)
}

func (service *ConnectionsService) InsertUserConnection(connection *domain.UserConnection) error {
	return service.graph.InsertUserConnection(connection)
}

func (service *ConnectionsService) DeleteUserConnection(connection *domain.UserConnection) error {
	return service.graph.DeleteUserConnection(connection)
}

func (service *ConnectionsService) InsertConnectionRequest(connection *domain.UserConnection) error {
	return service.graph.InsertConnectionRequest(connection)
}
func (service *ConnectionsService) CancelConnectionRequest(connection *domain.UserConnection) error {
	return service.graph.CancelConnectionRequest(connection)
}

func (service *ConnectionsService) AcceptUserConnection(connection *domain.UserConnection) error {
	return service.graph.AcceptUserConnectionRequest(connection)
}
func (service *ConnectionsService) DeclineUserConnection(connection *domain.UserConnection) error {
	return service.graph.DeclineUserConnectionRequest(connection)
}

func (service *ConnectionsService) BlockUser(connection *domain.UserConnection) error {
	return service.graph.BlockUser(connection)
}
func (service *ConnectionsService) UnblockUser(connection *domain.UserConnection) error {
	return service.graph.UnblockUser(connection)
}

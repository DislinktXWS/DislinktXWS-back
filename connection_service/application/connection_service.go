package application

import (
	"context"
	"github.com/dislinktxws-back/connection_service/domain"
)

type ConnectionsService struct {
	graph domain.ConnectionsGraph
}

func NewConnectionsService(graph domain.ConnectionsGraph) *ConnectionsService {
	return &ConnectionsService{
		graph: graph,
	}
}

func (service *ConnectionsService) GetAll(userId string, ctx context.Context) []string {
	return service.graph.GetAll(userId, ctx)
}
func (service *ConnectionsService) GetAllConnectionRequests(userId string, ctx context.Context) []string {
	return service.graph.GetConnectionRequests(userId, ctx)
}
func (service *ConnectionsService) GetUserRecommendations(userId string, ctx context.Context) []string {
	return service.graph.GetUserRecommendations(userId, ctx)
}
func (service *ConnectionsService) GetConnectionStatus(user1 string, user2 string, ctx context.Context) string {
	return service.graph.GetConnectionStatus(user1, user2, ctx)
}
func (service *ConnectionsService) GetBlockedUsers(userId string, ctx context.Context) []string {
	return service.graph.GetBlockedUsers(userId, ctx)
}

func (service *ConnectionsService) InsertNewUser(user string) error {
	return service.graph.InsertNewUser(user)
}

func (service *ConnectionsService) InsertUserConnection(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.InsertUserConnection(connection, ctx)
}

func (service *ConnectionsService) DeleteUserConnection(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.DeleteUserConnection(connection, ctx)
}

func (service *ConnectionsService) InsertConnectionRequest(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.InsertConnectionRequest(connection, ctx)
}
func (service *ConnectionsService) CancelConnectionRequest(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.CancelConnectionRequest(connection, ctx)
}

func (service *ConnectionsService) AcceptUserConnection(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.AcceptUserConnectionRequest(connection, ctx)
}
func (service *ConnectionsService) DeclineUserConnection(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.DeclineUserConnectionRequest(connection, ctx)
}

func (service *ConnectionsService) BlockUser(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.BlockUser(connection, ctx)
}
func (service *ConnectionsService) UnblockUser(connection *domain.UserConnection, ctx context.Context) error {
	return service.graph.UnblockUser(connection, ctx)
}

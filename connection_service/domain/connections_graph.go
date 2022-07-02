package domain

import "context"

type ConnectionsGraph interface {
	GetAll(username string, ctx context.Context) []string
	GetBlockedUsers(username string, ctx context.Context) []string
	GetConnectionRequests(username string, ctx context.Context) []string

	InsertNewUser(user string) error
	GetConnectionStatus(user1 string, user2 string, ctx context.Context) string

	DeleteUserConnection(connection *UserConnection, ctx context.Context) error
	InsertUserConnection(connection *UserConnection, ctx context.Context) error

	InsertConnectionRequest(connection *UserConnection, ctx context.Context) error
	CancelConnectionRequest(connection *UserConnection, ctx context.Context) error
	AcceptUserConnectionRequest(connection *UserConnection, ctx context.Context) error
	DeclineUserConnectionRequest(connection *UserConnection, ctx context.Context) error

	BlockUser(connection *UserConnection, ctx context.Context) error
	UnblockUser(connection *UserConnection, ctx context.Context) error

	GetUserRecommendations(user string, ctx context.Context) []string
}

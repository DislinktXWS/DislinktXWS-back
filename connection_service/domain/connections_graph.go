package domain

type ConnectionsGraph interface {
	//Get(connection *UserConnection) bool
	GetAll(username string) []string
	GetBlockedUsers(username string) []string
	GetConnectionRequests(username string) []string

	InsertNewUser(user string) error
	GetConnectionStatus(user1 string, user2 string) string

	DeleteUserConnection(connection *UserConnection) error
	InsertUserConnection(connection *UserConnection) error

	InsertUserConnectionRequest(connection *UserConnection) error
	AcceptUserConnectionRequest(connection *UserConnection) error
	DeclineUserConnectionRequest(connection *UserConnection) error

	BlockUser(connection *UserConnection) error
	UnblockUser(connection *UserConnection) error
}

package domain

type ConnectionsGraph interface {
	Get(connection *UserConnection) bool
	GetAll(username string) []string
	InsertNewUser(user string) error
	InsertUserConnection(connection *UserConnection) error
	DeleteUserConnection(connection *UserConnection) error
}

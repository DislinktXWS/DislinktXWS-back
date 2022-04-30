package domain

type ConnectionsGraph interface {
	Get(connection *UserConnection) bool
	GetAll(username string) []*User
	InsertNewUser(user *User) error
	InsertUserConnection(connection *UserConnection) error
	DeleteUserConnection(connection *UserConnection) error
}

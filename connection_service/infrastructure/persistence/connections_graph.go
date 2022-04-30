package persistence

import (
	"module/connection_service/domain"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type ConnectionsDBGraph struct {
	session *neo4j.Session
}

func NewConnectionsGraph(session *neo4j.Session) domain.ConnectionsGraph {
	return &ConnectionsDBGraph{
		session: session,
	}
}

func (store *ConnectionsDBGraph) Get(connection *domain.UserConnection) bool {
	//filter := bson.M{"_id": id} //M je getovanje ali NE po redosledu kakav je u bazi
	return true
}

func (store *ConnectionsDBGraph) GetAll(user string) []string {
	var session = *store.session
	connections, _ := getUserConnections(session, user)
	return connections
}

func (store *ConnectionsDBGraph) InsertNewUser(user string) error {

	var session = *store.session
	_, err := session.WriteTransaction(addUserNodeTxFunc(user))
	return err
}

func (store *ConnectionsDBGraph) InsertUserConnection(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(connectUsersTxFunc(connection.Connecting, connection.Connected))
	return err
}

func (store *ConnectionsDBGraph) DeleteUserConnection(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(disconnectUsersTxFunc(connection.Connecting, connection.Connected))
	return err
}

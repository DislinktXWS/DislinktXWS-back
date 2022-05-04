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

func (store *ConnectionsDBGraph) GetBlockedUsers(user string) []string {
	var session = *store.session
	blocked, _ := getBlockedUsersTxFunc(session, user)
	return blocked
}

func (store *ConnectionsDBGraph) GetConnectionRequests(user string) []string {
	var session = *store.session
	blocked, _ := getConnectionRequestsTxFunc(session, user)
	return blocked
}

func (store *ConnectionsDBGraph) InsertNewUser(user string) error {
	var session = *store.session
	_, err := session.WriteTransaction(addUserNodeTxFunc(user))
	return err
}

func (store *ConnectionsDBGraph) InsertUserConnection(connection *domain.UserConnection) error {
	var session = *store.session
	_, e1 := session.WriteTransaction(connectUsersTxFunc(connection.Connecting, connection.Connected))
	if e1 != nil {
		return e1
	}
	_, e2 := session.WriteTransaction(connectUsersTxFunc(connection.Connected, connection.Connecting))
	return e2
}

func (store *ConnectionsDBGraph) DeleteUserConnection(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(disconnectUsersTxFunc(connection.Connecting, connection.Connected))
	return err
}

func (store *ConnectionsDBGraph) InsertUserConnectionRequest(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(requestConnectionTxFunc(connection.Connecting, connection.Connected))
	return err
}

func (store *ConnectionsDBGraph) AcceptUserConnectionRequest(connection *domain.UserConnection) error {
	var session = *store.session
	_, er := session.WriteTransaction(deleteConnectionRequestTxFunc(connection.Connecting, connection.Connected))
	if er != nil {
		return er
	}
	_, err := session.WriteTransaction(connectUsersTxFunc(connection.Connecting, connection.Connected))
	if err != nil {
		return err
	}
	_, e := session.WriteTransaction(connectUsersTxFunc(connection.Connected, connection.Connecting))
	return e
}
func (store *ConnectionsDBGraph) DeclineUserConnectionRequest(connection *domain.UserConnection) error {
	var session = *store.session
	_, er := session.WriteTransaction(deleteConnectionRequestTxFunc(connection.Connecting, connection.Connected))
	return er
}
func (store *ConnectionsDBGraph) BlockUser(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(blockUserTxFunc(connection.Connecting, connection.Connected))
	if err != nil {
		return err
	}
	_, er := session.WriteTransaction(disconnectUsersTxFunc(connection.Connecting, connection.Connected))
	if er != nil {
		return er
	}
	_, error := session.WriteTransaction(disconnectUsersTxFunc(connection.Connected, connection.Connecting))

	return error
}
func (store *ConnectionsDBGraph) UnblockUser(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(unblockUserTxFunc(connection.Connecting, connection.Connected))
	return err
}

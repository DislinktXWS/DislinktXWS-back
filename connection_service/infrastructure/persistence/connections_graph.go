package persistence

import (
	"github.com/dislinktxws-back/connection_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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

func (store *ConnectionsDBGraph) GetConnectionStatus(user1 string, user2 string) string {
	var session = *store.session

	connected, _ := checkIfConnectionExists(session, user1, user2, "CONNECTED")
	if connected {
		return "connected"
	}
	blockedByYou, _ := checkIfConnectionExists(session, user1, user2, "BLOCKED")
	if blockedByYou {
		return "blockedByYou"
	}
	blockedYou, _ := checkIfConnectionExists(session, user2, user1, "BLOCKED")
	if blockedYou {
		return "blockedYou"
	}
	connectionRequestedByYou, _ := checkIfConnectionExists(session, user1, user2, "REQUESTED_CONNECTION")
	if connectionRequestedByYou {
		return "connectionRequestedByYou"
	}
	connectionRequestedByUser, _ := checkIfConnectionExists(session, user2, user1, "REQUESTED_CONNECTION")
	if connectionRequestedByUser {
		return "connectionRequestedByUser"
	}

	return "none"
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

func (store *ConnectionsDBGraph) GetUserRecommendations(user string) []string {
	var session = *store.session
	friendsOfFriends, _ := getFriendsOfFriendsTxFunc(session, user)
	randomUsers, _ := getRandomUsersTxFunc(session, user)

	if len(friendsOfFriends) < 10 {

		for _, v := range randomUsers {
			if !contains(friendsOfFriends, v) {
				friendsOfFriends = append(friendsOfFriends, v)
			}
			if len(friendsOfFriends) == 10 {
				break
			}
		}
	}
	return friendsOfFriends
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
	_, e1 := session.WriteTransaction(disconnectUsersTxFunc(connection.Connecting, connection.Connected))
	if e1 != nil {
		return e1
	}
	_, e2 := session.WriteTransaction(disconnectUsersTxFunc(connection.Connected, connection.Connecting))
	return e2
}

func (store *ConnectionsDBGraph) InsertConnectionRequest(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(requestConnectionTxFunc(connection.Connecting, connection.Connected))
	return err
}

func (store *ConnectionsDBGraph) CancelConnectionRequest(connection *domain.UserConnection) error {
	var session = *store.session
	_, err := session.WriteTransaction(cancelConnectionRequestTxFunc(connection.Connecting, connection.Connected))
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

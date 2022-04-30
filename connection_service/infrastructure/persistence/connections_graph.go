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

func (store *ConnectionsDBGraph) GetAll(username string) []string {
	//filter := bson.D{{}} //D je getovanje ali  po redosledu kakav je u bazi
	return nil
}

func (store *ConnectionsDBGraph) InsertNewUser(user string) error {

	var session = *store.session
	_, err := session.WriteTransaction(addUserNodeTxFunc(user))
	return err
}

func addUserNodeTxFunc(user string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("CREATE (a:User {id: $name})", map[string]interface{}{"name": user})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	}
}

func (store *ConnectionsDBGraph) InsertUserConnection(connection *domain.UserConnection) error {
	return nil
}
func (store *ConnectionsDBGraph) DeleteUserConnection(connection *domain.UserConnection) error {
	return nil
}

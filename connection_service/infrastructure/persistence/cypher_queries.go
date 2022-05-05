package persistence

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func addUserNodeTxFunc(user string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("CREATE (a:User {userId: $user})", map[string]interface{}{"user": user})
		if err != nil {
			return nil, err
		}
		return result.Consume()
	}
}

func checkIfConnectionExists(session neo4j.Session, user1 string, user2 string, connection string) (bool, error) {

	ret, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var exists bool

		result, err := tx.Run("MATCH (user: User {userId: $user1})-[rel:$connection]->(userConnection:User{userId: $user2}) RETURN rel",
			map[string]interface{}{"user1": user1, "user2": user2, "connection": connection})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			exists = true
		} else {
			exists = false
		}

		return exists, nil
	})
	if err != nil {
		return false, err
	}
	return ret.(bool), nil
}

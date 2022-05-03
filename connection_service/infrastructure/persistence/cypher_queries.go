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

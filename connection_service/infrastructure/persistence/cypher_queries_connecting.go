package persistence

import "github.com/neo4j/neo4j-go-driver/neo4j"

func connectUsersTxFunc(connecting string, connected string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"MATCH (user:User {userId: $connecting}) "+
				"MATCH (secondUser:User {userId: $connected}) "+
				"CREATE (user)-[:CONNECTED]->(secondUser)", map[string]interface{}{"connecting": connecting, "connected": connected})

		if err != nil {
			return nil, err
		}

		return result.Consume()
	}
}

func disconnectUsersTxFunc(disconnecting string, disconnected string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"MATCH (user:User {userId: $disconnecting}) "+
				"- [rel:CONNECTED] -> (secondUser:User {userId: $disconnected}) "+
				"DELETE rel", map[string]interface{}{"disconnecting": disconnecting, "disconnected": disconnected})

		if err != nil {
			return nil, err
		}
		return result.Consume()
	}
}

func getUserConnections(session neo4j.Session, user string) ([]string, error) {

	people, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		result, err := tx.Run("MATCH (user: User {userId: $user})-[:CONNECTED]->(userConnection:User) RETURN userConnection.userId", map[string]interface{}{"user": user})
		if err != nil {
			return nil, err
		}

		for result.Next() {
			list = append(list, result.Record().GetByIndex(0).(string))
		}

		if err = result.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return people.([]string), nil
}

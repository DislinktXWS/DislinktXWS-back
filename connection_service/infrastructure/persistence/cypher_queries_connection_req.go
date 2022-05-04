package persistence

import "github.com/neo4j/neo4j-go-driver/neo4j"

func requestConnectionTxFunc(connecting string, connected string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"MATCH (user:User {userId: $connecting}) "+
				"MATCH (secondUser:User {userId: $connected}) "+
				"CREATE (user)-[:REQUESTED_CONNECTION]->(secondUser)", map[string]interface{}{"connecting": connecting, "connected": connected})

		if err != nil {
			return nil, err
		}

		return result.Consume()
	}
}

func deleteConnectionRequestTxFunc(connecting string, connected string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"MATCH (user:User {userId: $connecting}) "+
				"- [rel:REQUESTED_CONNECTION] -> (secondUser:User {userId: $connected}) "+
				"DELETE rel", map[string]interface{}{"connecting": connecting, "connected": connected})

		if err != nil {
			return nil, err
		}
		return result.Consume()
	}
}

func getConnectionRequestsTxFunc(session neo4j.Session, user string) ([]string, error) {

	people, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		result, err := tx.Run("MATCH (user: User {userId: $user})<-[:REQUESTED_CONNECTION]-(requested:User) RETURN requested.userId", map[string]interface{}{"user": user})
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

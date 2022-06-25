package persistence

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

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

		result, err := tx.Run("MATCH (user: User {userId: $user1})-[rel:"+connection+"]->(userConnection:User{userId: $user2}) RETURN rel",
			map[string]interface{}{"user1": user1, "user2": user2})
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

func getFriendsOfFriendsTxFunc(session neo4j.Session, user string) ([]string, error) {

	people, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		result, err := tx.Run(
			"MATCH (u1: User)-[:CONNECTED]->(u2:User)<-[:CONNECTED]-(u3:User) "+
				"WHERE u1.userId=$user AND u3.userId<>$user "+
				"AND NOT exists((u1)-[:CONNECTED]-(u3)) "+
				"AND NOT exists((u1)-[:BLOCKED]-(u3)) "+
				"RETURN distinct u3.userId LIMIT 10 ",

			map[string]interface{}{"user": user})
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

func getRandomUsersTxFunc(session neo4j.Session, user string) ([]string, error) {

	people, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		result, err := tx.Run(
			"MATCH (user: User { userId: $user}),(N:User) "+
				"WHERE NOT exists((N)-[:CONNECTED]-(user)) "+
				"AND NOT exists((N)-[:BLOCKED]-(user)) "+
				"RETURN N LIMIT 20 ",

			map[string]interface{}{"user": user})
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

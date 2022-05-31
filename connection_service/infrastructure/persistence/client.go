package persistence

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

//vraca sesiju

func GetClient(username, password, uri string) (*neo4j.Session, error) {

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	//defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	//defer session.Close()

	if err == nil {
		fmt.Print("connection established")
	}
	return &session, err
}

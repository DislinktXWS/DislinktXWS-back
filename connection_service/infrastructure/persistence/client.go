package persistence

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

//vraca sesiju

func GetClient(username, password, uri string) (*neo4j.Session, error) {

	driver, _ := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), func(c *neo4j.Config) { c.Encrypted = false })
	defer driver.Close()

	session, err := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	return &session, err
}

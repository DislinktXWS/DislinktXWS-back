package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host                     string
	Port                     string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	InsertUserCommandSubject string
	InsertUserReplySubject   string

	Uri      string
	Username string
	Password string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:                     os.Getenv("CONNECTION_SERVICE_HOST"),
		Port:                     os.Getenv("CONNECTION_SERVICE_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		InsertUserCommandSubject: os.Getenv("INSERT_USER_COMMAND_SUBJECT"),
		InsertUserReplySubject:   os.Getenv("INSERT_USER_REPLY_SUBJECT"),

		Uri:      "neo4j://neo4j:7687",
		Username: "neo4j",
		Password: "ConnectionDB", //password je onaj koji je postavi kad se pravi nova sema bp
	}
}

//var uri string = "bolt://localhost:7687"
//var username string = "neo4j"
//var password string = "ConnectionDB"

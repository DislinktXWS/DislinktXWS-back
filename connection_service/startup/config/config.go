package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host string
	Port string

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
		Host: os.Getenv("CONNECTION_SERVICE_HOST"),
		Port: os.Getenv("CONNECTION_SERVICE_PORT"),

		//Uri:      "bullshit",
		Uri:      "neo4j://neo4j:7687",
		Username: os.Getenv("CONNECTION_DB_USER"),
		Password: os.Getenv("CONNECTION_DB_PASS"), //password je onaj koji je postavi kad se pravi nova sema bp
	}
}

//var uri string = "bolt://localhost:7687"
//var username string = "neo4j"
//var password string = "ConnectionDB"

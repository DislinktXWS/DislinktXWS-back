package config

type Config struct {
	Host string
	Port string

	Uri      string
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: "8089",

		Uri:      "bolt://localhost:7687",
		Username: "neo4j",
		Password: "ConnectionDB", //password je onaj koji je postavi kad se pravi nova sema bp
	}
}

//var uri string = "bolt://localhost:7687"
//var username string = "neo4j"
//var password string = "ConnectionDB"

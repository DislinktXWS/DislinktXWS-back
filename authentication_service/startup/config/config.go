package config

import (
	"os"
)

type Config struct {
	Host         string
	Port         string
	PostDBHost   string
	PostDBPort   string
	JWTSecretKey string
}

func NewConfig() *Config {
	/*devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}*/
	return &Config{
		Host:         os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		Port:         os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		PostDBHost:   os.Getenv("AUTH_DB_HOST"),
		PostDBPort:   os.Getenv("AUTH_DB_PORT"),
		JWTSecretKey: "r43t18sc",
	}
}

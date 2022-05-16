package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host       string
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:       os.Getenv("POST_SERVICE_HOST"),
		Port:       os.Getenv("POST_SERVICE_PORT"),
		PostDBHost: os.Getenv("POST_DB_HOST"),
		PostDBPort: os.Getenv("POST_DB_PORT"),
	}
}

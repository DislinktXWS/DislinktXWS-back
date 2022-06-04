package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host          string
	Port          string
	MessageDBHost string
	MessageDBPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:          os.Getenv("MESSAGE_SERVICE_HOST"),
		Port:          os.Getenv("MESSAGE_SERVICE_PORT"),
		MessageDBHost: os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort: os.Getenv("MESSAGE_DB_PORT"),
	}
}

package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host       string
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:       os.Getenv("USER_SERVICE_HOST"),
		Port:       os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: os.Getenv("USER_DB_HOST"),
		UserDBPort: os.Getenv("USER_DB_PORT"),
	}
}

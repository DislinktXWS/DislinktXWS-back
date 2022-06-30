package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host               string
	Port               string
	NotificationDBHost string
	NotificationDBPort string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}

	return &Config{
		Host:               os.Getenv("NOTIFICATIONS_SERVICE_HOST"),
		Port:               os.Getenv("NOTIFICATIONS_SERVICE_PORT"),
		NotificationDBHost: os.Getenv("NOTIFICATIONS_DB_HOST"),
		NotificationDBPort: os.Getenv("NOTIFICATIONS_DB_PORT"),
	}
}

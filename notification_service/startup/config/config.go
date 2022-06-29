package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
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
		Host:               "",
		Port:               "",
		NotificationDBHost: "",
		NotificationDBPort: "",
	}
}

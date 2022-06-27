package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct { //prijava svih servisa gateway-u
	Host               string
	Port               string
	UserHost           string
	UserPort           string
	PostHost           string
	PostPort           string
	ConnectionHost     string
	ConnectionPort     string
	AuthenticationHost string
	AuthenticationPort string
	MessageHost        string
	MessagePort        string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:               os.Getenv("GATEWAY_HOST"),
		Port:               os.Getenv("GATEWAY_PORT"),
		UserHost:           os.Getenv("USER_SERVICE_HOST"),
		UserPort:           os.Getenv("USER_SERVICE_PORT"),
		PostHost:           os.Getenv("POST_SERVICE_HOST"),
		PostPort:           os.Getenv("POST_SERVICE_PORT"),
		ConnectionHost:     os.Getenv("CONNECTION_SERVICE_HOST"),
		ConnectionPort:     os.Getenv("CONNECTION_SERVICE_PORT"),
		AuthenticationHost: os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort: os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		MessageHost:        os.Getenv("MESSAGE_SERVICE_HOST"),
		MessagePort:        os.Getenv("MESSAGE_SERVICE_PORT"),
	}
}

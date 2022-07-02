package config

import (
	"flag"
	"os"

	cnf "github.com/dislinktxws-back/common/config"
)

type Config struct { //prijava svih servisa gateway-u
	Host                   string
	Port                   string
	UserHost               string
	UserPort               string
	PostHost               string
	PostPort               string
	ConnectionHost         string
	ConnectionPort         string
	AuthenticationHost     string
	AuthenticationPort     string
	BusinessOfferHost      string
	BusinessOfferPort      string
	NotificationsOfferHost string
	NotificationsOfferPort string
	MessageHost            string
	MessagePort            string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:                   os.Getenv("GATEWAY_HOST"),
		Port:                   os.Getenv("GATEWAY_PORT"),
		UserHost:               os.Getenv("USER_SERVICE_HOST"),
		UserPort:               os.Getenv("USER_SERVICE_PORT"),
		PostHost:               os.Getenv("POST_SERVICE_HOST"),
		PostPort:               os.Getenv("POST_SERVICE_PORT"),
		ConnectionHost:         os.Getenv("CONNECTION_SERVICE_HOST"),
		ConnectionPort:         os.Getenv("CONNECTION_SERVICE_PORT"),
		AuthenticationHost:     os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		AuthenticationPort:     os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		BusinessOfferHost:      os.Getenv("BUSINESS_OFFER_SERVICE_HOST"),
		BusinessOfferPort:      os.Getenv("BUSINESS_OFFER_SERVICE_PORT"),
		NotificationsOfferHost: os.Getenv("NOTIFICATIONS_SERVICE_HOST"),
		NotificationsOfferPort: os.Getenv("NOTIFICATIONS_SERVICE_PORT"),
		MessageHost:            os.Getenv("MESSAGE_SERVICE_HOST"),
		MessagePort:            os.Getenv("MESSAGE_SERVICE_PORT"),
	}
}

package config

import (
	"os"
)

type Config struct {
	Host     string
	Port     string
	Uri      string
	Username string
	Password string
}

func NewConfig() *Config {

	return &Config{
		Host:     os.Getenv("BUSINESS_OFFER_SERVICE_HOST"),
		Port:     os.Getenv("BUSINESS_OFFER_SERVICE_PORT"),
		Uri:      "neo4j://neo4j:7687",
		Username: "neo4j",
		Password: "BusinessOfferDB",
	}
}

package config

type Config struct { //prijava svih servisa gateway-u
	Host           string
	Port           string
	UserHost       string
	UserPort       string
	PostHost       string
	PostPort       string
	ConnectionHost string
	ConnectionPort string
}

func NewConfig() *Config {
	return &Config{
		Host:           "localhost",
		Port:           "8083",
		UserHost:       "localhost",
		UserPort:       "8081",
		PostHost:       "localhost",
		PostPort:       "8085",
		ConnectionHost: "localhost",
		ConnectionPort: "8087",
	}
}

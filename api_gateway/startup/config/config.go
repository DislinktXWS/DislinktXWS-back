package config

type Config struct { //prijava svih servisa gateway-u
	Host     string
	Port     string
	UserHost string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "8080",
		UserHost: "localhost",
		UserPort: "8081",
	}
}

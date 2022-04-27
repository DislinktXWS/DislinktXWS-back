package config

type Config struct {
	Host       string
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Host:       "localhost",
		Port:       "8081",
		UserDBHost: "localhost",
		UserDBPort: "27017",
	}
}

package config

type Config struct {
	Host       string
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	return &Config{
		Host:       "localhost",
		Port:       "8085",
		PostDBHost: "localhost",
		PostDBPort: "27017",
	}
}

package config

type Config struct {
	Host         string
	Port         string
	PostDBHost   string
	PostDBPort   string
	JWTSecretKey string
}

func NewConfig() *Config {
	return &Config{
		Host:         "localhost",
		Port:         "8083",
		PostDBHost:   "localhost",
		PostDBPort:   "27017",
		JWTSecretKey: "r43t18sc",
	}
}

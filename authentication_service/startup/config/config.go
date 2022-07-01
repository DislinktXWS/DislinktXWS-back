package config

import (
	"os"
)

type Config struct {
	Host                     string
	Port                     string
	PostDBHost               string
	PostDBPort               string
	JWTSecretKey             string
	TwoFactorDBHost          string
	TwoFactorDBPort          string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	InsertUserCommandSubject string
	InsertUserReplySubject   string
}

func NewConfig() *Config {
	/*devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()
	if *devEnv {
		cnf.LoadEnv()
	}*/
	return &Config{
		Host:                     os.Getenv("AUTHENTICATION_SERVICE_HOST"),
		Port:                     os.Getenv("AUTHENTICATION_SERVICE_PORT"),
		PostDBHost:               os.Getenv("AUTH_DB_HOST"),
		PostDBPort:               os.Getenv("AUTH_DB_PORT"),
		JWTSecretKey:             "r43t18sc",
		TwoFactorDBHost:          os.Getenv("TWOFACTOR_DB_HOST"),
		TwoFactorDBPort:          os.Getenv("TWOFACTOR_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		InsertUserCommandSubject: os.Getenv("INSERT_USER_COMMAND_SUBJECT"),
		InsertUserReplySubject:   os.Getenv("INSERT_USER_REPLY_SUBJECT"),
	}
}

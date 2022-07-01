package config

import (
	"flag"
	cnf "github.com/dislinktxws-back/common/config"
	"os"
)

type Config struct {
	Host                     string
	Port                     string
	UserDBHost               string
	UserDBPort               string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	InsertUserCommandSubject string
	InsertUserReplySubject   string
}

func NewConfig() *Config {
	devEnv := flag.Bool("dev", false, "use dev environment variables")
	flag.Parse()

	if *devEnv {
		cnf.LoadEnv()
	}
	return &Config{
		Host:                     os.Getenv("USER_SERVICE_HOST"),
		Port:                     os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:               os.Getenv("USER_DB_HOST"),
		UserDBPort:               os.Getenv("USER_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		InsertUserCommandSubject: os.Getenv("INSERT_USER_COMMAND_SUBJECT"),
		InsertUserReplySubject:   os.Getenv("INSERT_USER_REPLY_SUBJECT"),
	}
}

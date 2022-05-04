package main

import (
	"module/authentication_service/startup"
	"module/authentication_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

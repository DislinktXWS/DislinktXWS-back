package main

import (
	"module/connection_service/startup"
	"module/connection_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

package main

import (
	"module/user_service/startup"
	"module/user_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

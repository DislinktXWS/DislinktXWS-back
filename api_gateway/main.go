package main

import (
	"module/api_gateway/startup"
	"module/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

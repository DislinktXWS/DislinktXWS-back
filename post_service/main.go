package main

import (
	"module/post_service/startup"
	"module/post_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

package main

import (
	"github.com/dislinktxws-back/connection_service/startup"
	"github.com/dislinktxws-back/connection_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

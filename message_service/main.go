package main

import (
	"github.com/dislinktxws-back/message_service/startup"
	"github.com/dislinktxws-back/message_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

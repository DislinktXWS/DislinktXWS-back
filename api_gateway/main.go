package main

import (
	"github.com/dislinktxws-back/api_gateway/startup"
	"github.com/dislinktxws-back/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

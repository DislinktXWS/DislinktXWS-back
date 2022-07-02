package main

import (
	"github.com/dislinktxws-back/business_offer_service/startup"
	config2 "github.com/dislinktxws-back/business_offer_service/startup/config"
)

func main() {
	config := config2.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}

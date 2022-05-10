package api

import (
	"github.com/dislinktxws-back/connection_service/application"
)

type CreateConnectionCommandHandler struct {
	connectionService *application.ConnectionsService
}

func NewCreateConnectionCommandHandler(connectionsService *application.ConnectionsService) (*CreateConnectionCommandHandler, error) {
	o := &CreateConnectionCommandHandler{
		connectionService: connectionsService,
	}
	return o, nil
}

package api

import (
	"module/connection_service/application"
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

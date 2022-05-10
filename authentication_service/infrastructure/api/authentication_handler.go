package api

import (
	"github.com/dislinktxws-back/authentication_service/application"
)

type CreateAuthCommandHandler struct {
	authenticationService *application.AuthenticationService
}

func NewCreateUserCommandHandler(authenticationService *application.AuthenticationService) (*CreateAuthCommandHandler, error) {
	o := &CreateAuthCommandHandler{
		authenticationService: authenticationService,
	}
	return o, nil
}

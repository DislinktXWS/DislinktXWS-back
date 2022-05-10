package api

import (
	"github.com/dislinktxws-back/user_service/application"
)

type CreateUserCommandHandler struct {
	userService *application.UserService
}

func NewCreateUserCommandHandler(userService *application.UserService) (*CreateUserCommandHandler, error) {
	o := &CreateUserCommandHandler{
		userService: userService,
	}
	return o, nil
}

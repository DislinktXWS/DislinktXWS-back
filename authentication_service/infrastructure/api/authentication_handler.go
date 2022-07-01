package api

import (
	"github.com/dislinktxws-back/authentication_service/application"
	events "github.com/dislinktxws-back/common/saga/insert_user"
	saga "github.com/dislinktxws-back/common/saga/messaging"
)

type CreateAuthCommandHandler struct {
	authenticationService *application.AuthenticationService
	replyPublisher        saga.Publisher
	commandSubscriber     saga.Subscriber
}

func NewCreateAuthCommandHandler(authenticationService *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateAuthCommandHandler, error) {
	o := &CreateAuthCommandHandler{
		authenticationService: authenticationService,
		replyPublisher:        publisher,
		commandSubscriber:     subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateAuthCommandHandler) handle(command *events.InsertUserCommand) {
	reply := events.InsertUserReply{User: command.User}

	switch command.Type {
	case events.InsertUserAuthentication:
		user := mapCommandToAuth(command)
		err := handler.authenticationService.Register(user)
		if err != nil {
			reply.Type = events.UserAuthenticationNotInserted
			break
		}
		reply.Type = events.UserAuthenticationInserted
	case events.RollbackInsertUserAuthentication:
		err := handler.authenticationService.Delete(command.User.Id)
		if err != nil {
			return
		}
		reply.Type = events.UserAuthenticationRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}

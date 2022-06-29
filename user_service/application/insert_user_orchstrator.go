package application

import (
	events "github.com/dislinktxws-back/common/saga/insert_user"
	saga "github.com/dislinktxws-back/common/saga/messaging"
	"github.com/dislinktxws-back/user_service/domain"
)

type InsertUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewInsertUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*InsertUserOrchestrator, error) {
	o := &InsertUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *InsertUserOrchestrator) Start(user *domain.User) error {
	event := &events.InsertUserCommand{
		Type: events.InsertUser,
		User: events.User{
			Id:                user.Id.Hex(),
			Name:              user.Name,
			Surname:           user.Surname,
			Username:          user.Username,
			DateOfBirth:       user.DateOfBirth,
			Gender:            user.Gender,
			Email:             user.Email,
			Phone:             user.Phone,
			Biography:         user.Biography,
			IsPublic:          true,
			VerificationToken: "",
			ApiKey:            "",
		},
	}
	return o.commandPublisher.Publish(event)
}

func (o *InsertUserOrchestrator) handle(reply *events.InsertUserReply) {
	command := events.InsertUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *InsertUserOrchestrator) nextCommandType(reply events.InsertUserReplyType) events.InsertUserCommandType {
	switch reply {
	case events.UserInserted:
		return events.InsertUserNode
	case events.UserNotInserted:
		return events.RollbackInsertUser
	case events.UserNodeInserted:
		return events.InsertUserNode
	case events.UserNodeNotInserted:
		return events.RollbackInsertUserNode
	default:
		return events.UnknownCommand
	}
}

package application

import (
	"fmt"
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
	fmt.Println("STARTUJEMO ORKESTRATOR")
	fmt.Println(user.Id)
	fmt.Println(user.VerificationToken)
	event := &events.InsertUserCommand{
		Type: events.InsertUserAuthentication,
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
			VerificationToken: user.VerificationToken,
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
	case events.UserAuthenticationInserted:
		fmt.Println("REPLY: 0")
		return events.InsertUserNode
	case events.UserAuthenticationNotInserted:
		fmt.Println("REPLY: 1")
		return events.RollbackInsertUser
	case events.UserAuthenticationRolledBack:
		fmt.Println("REPLY: 2")
		return events.RollbackInsertUser
	case events.UserNodeInserted:
		fmt.Println("REPLY: 3")
		return events.UnknownCommand //successful end
	case events.UserNodeNotInserted:
		fmt.Println("REPLY: 4")
		return events.RollbackInsertUserAuthentication
	case events.UserInsertRolledBack:
		fmt.Println("REPLY: 5")
		return events.UnknownCommand //unsuccessful end
	default:
		fmt.Println("NEPOZNATI REPLY")
		return events.UnknownCommand
	}
}

package api

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/authentication_service/application"
	"github.com/dislinktxws-back/authentication_service/utils"
	pb "github.com/dislinktxws-back/common/proto/authentication_service"
	events "github.com/dislinktxws-back/common/saga/insert_user"
	saga "github.com/dislinktxws-back/common/saga/messaging"
	"log"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	service           *application.AuthenticationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewAuthenticationHandler(service *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) *AuthenticationHandler {
	o := &AuthenticationHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	o.commandSubscriber.Subscribe(o.handle)
	return o
}

func init() {
	infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(infoFile, "INFO: ", log.LstdFlags|log.Lshortfile)

	errFile, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	ErrorLogger = log.New(errFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func (handler *AuthenticationHandler) handle(command *events.InsertUserCommand) {
	reply := events.InsertUserReply{User: command.User}
	fmt.Println("AUTH HANDLER")
	fmt.Println(command.Type)

	switch command.Type {
	case events.InsertUserAuthentication:
		fmt.Println("REGISTRUJEMO USERA")
		user := mapCommandToAuth(command)
		err := handler.service.Register(user)
		if err != nil {
			reply.Type = events.UserAuthenticationNotInserted
			break
		}
		reply.Type = events.UserAuthenticationInserted
	case events.RollbackInsertUserAuthentication:
		err := handler.service.Delete(command.User.Id)
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

func (handler *AuthenticationHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	newAuth := mapAuth(request.Auth)
	status, err, token, isTwoFactorEnabled := handler.service.Login(newAuth)
	if status != 200 {
		ErrorLogger.Println("Action: 28, Message: Wrong credentials for login")
	}
	if status == 200 {
		InfoLogger.Println("Action: 29, Message: Login successfull")
	}
	return &pb.LoginResponse{
		Status:             status,
		Error:              err,
		Token:              token,
		IsTwoFactorEnabled: isTwoFactorEnabled,
	}, nil
}

func (handler *AuthenticationHandler) PasswordlessLogin(ctx context.Context, request *pb.PasswordlessLoginRequest) (*pb.PasswordlessLoginResponse, error) {
	status, err, token := handler.service.PasswordlessLogin(request.VerificationToken)
	if status != 200 {
		ErrorLogger.Println("Action: 28, Message: Wrong email for passwordless login")
	}
	if status == 200 {
		InfoLogger.Println("Action: 29, Message: Login successfull")
	}
	return &pb.PasswordlessLoginResponse{
		Status: status,
		Error:  err,
		Token:  token,
	}, nil
}

func (handler *AuthenticationHandler) Validate(ctx context.Context, request *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	token := request.Token
	status, err, user := handler.service.Validate(token)
	if status != 200 {
		ErrorLogger.Println("Action: 30, Message: Token is not valid or expired!")
	}
	if status == 200 {
		InfoLogger.Println("Action: 31, Message: User validation successfull")
	}
	return &pb.ValidateResponse{
		Status: status,
		Error:  err,
		User:   user,
	}, nil
}

func (handler *AuthenticationHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	Auth := mapAuth(request.Auth)
	fmt.Println(request.Auth)
	Auth.Password = utils.HashPassword(Auth.Password)
	Auth.IsVerified = false
	Auth.VerificationCreationTime = time.Now()
	err := handler.service.Register(Auth)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Can not register user!")
		return nil, err
	}
	InfoLogger.Println("Action: 3, Message: User " + Auth.Username + " registered successfully!")
	return &pb.RegisterResponse{}, nil
}

func (handler *AuthenticationHandler) EditUsername(ctx context.Context, request *pb.EditUsernameRequest) (*pb.EditUsernameResponse, error) {
	auth := mapAuth(request.Auth)
	_, err := handler.service.EditUsername(auth)
	if err != nil {
		ErrorLogger.Println("Action: 5, Message: Username is not unique!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User " + auth.Username + " edited successfully")
	return &pb.EditUsernameResponse{}, nil
}

func (handler *AuthenticationHandler) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	auth := mapAuth(request.Auth)
	err := handler.service.ChangePassword(auth)
	if err != nil {
		ErrorLogger.Println("Action 5, Message: Cannot change password!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User " + auth.Username + " changed password successfully")
	return &pb.ChangePasswordResponse{}, nil
}

func (handler *AuthenticationHandler) GenerateVerificationToken(ctx context.Context,
	request *pb.GenerateVerificationTokenRequest) (*pb.GenerateVerificationTokenResponse, error) {
	err := handler.service.GenerateVerificationToken(request.Email)
	InfoLogger.Println("Action: 32, Message: E-mail sent successfully")
	return &pb.GenerateVerificationTokenResponse{}, err
}

func (handler *AuthenticationHandler) AccountRecovery(ctx context.Context, request *pb.AccountRecoveryRequest) (*pb.AccountRecoveryResponse, error) {
	status, err := handler.service.AccountRecovery(request.Email)
	if status != 200 {
		ErrorLogger.Println("Action: 33, Message: Wrong email")
	}
	if status == 200 {
		InfoLogger.Println("Action: 32, Message: Recovery mail sent successfull")
	}
	return &pb.AccountRecoveryResponse{
		Status: status,
		Error:  err,
	}, nil
}

func (handler *AuthenticationHandler) ChangeTwoFactorAuth(ctx context.Context, request *pb.ChangeTwoFactorAuthRequest) (*pb.ChangeTwoFactorAuthResponse, error) {
	qrCode, err := handler.service.ChangeTwoFactorAuth(request.Username)
	return &pb.ChangeTwoFactorAuthResponse{
		QrCode: qrCode,
		Error:  err,
	}, nil
}

func (handler *AuthenticationHandler) GetTwoFactorAuth(ctx context.Context, request *pb.GetTwoFactorAuthRequest) (*pb.GetTwoFactorAuthResponse, error) {
	flag := handler.service.GetTwoFactorAuth(request.Username)
	return &pb.GetTwoFactorAuthResponse{IsEnabled: flag}, nil
}

func (handler *AuthenticationHandler) VerifyTwoFactorAuthToken(ctx context.Context, request *pb.VerifyTwoFactorAuthTokenRequest) (*pb.VerifyTwoFactorAuthTokenResponse, error) {
	status, err, token := handler.service.VerifyTwoFactorAuthToken(request.Username, request.Token)
	return &pb.VerifyTwoFactorAuthTokenResponse{
		Status: status,
		Error:  err,
		Token:  token,
	}, nil
}

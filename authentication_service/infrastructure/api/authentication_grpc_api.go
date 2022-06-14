package api

import (
	"context"
	"github.com/dislinktxws-back/authentication_service/application"
	"github.com/dislinktxws-back/authentication_service/utils"
	pb "github.com/dislinktxws-back/common/proto/authentication_service"
	"time"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	service *application.AuthenticationService
}

func NewAuthenticationHandler(service *application.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		service: service,
	}
}

func (handler *AuthenticationHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	newAuth := mapAuth(request.Auth)
	status, err, token := handler.service.Login(newAuth)
	return &pb.LoginResponse{
		Status: status,
		Error:  err,
		Token:  token,
	}, nil
}

func (handler *AuthenticationHandler) PasswordlessLogin(ctx context.Context, request *pb.PasswordlessLoginRequest) (*pb.PasswordlessLoginResponse, error) {
	status, err, token := handler.service.PasswordlessLogin(request.VerificationToken)
	return &pb.PasswordlessLoginResponse{
		Status: status,
		Error:  err,
		Token:  token,
	}, nil
}

func (handler *AuthenticationHandler) Validate(ctx context.Context, request *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	token := request.Token
	status, err, user := handler.service.Validate(token)
	return &pb.ValidateResponse{
		Status: status,
		Error:  err,
		User:   user,
	}, nil
}

func (handler *AuthenticationHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	Auth := mapAuth(request.Auth)
	Auth.Password = utils.HashPassword(Auth.Password)
	Auth.IsVerified = false
	Auth.VerificationCreationTime = time.Now()
	err := handler.service.Register(Auth)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (handler *AuthenticationHandler) EditUsername(ctx context.Context, request *pb.EditUsernameRequest) (*pb.EditUsernameResponse, error) {
	auth := mapAuth(request.Auth)
	_, err := handler.service.EditUsername(auth)
	if err != nil {
		return nil, err
	}
	return &pb.EditUsernameResponse{}, nil
}

func (handler *AuthenticationHandler) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	auth := mapAuth(request.Auth)
	err := handler.service.ChangePassword(auth)
	if err != nil {
		return nil, err
	}
	return &pb.ChangePasswordResponse{}, nil
}

func (handler *AuthenticationHandler) GenerateVerificationToken(ctx context.Context,
	request *pb.GenerateVerificationTokenRequest) (*pb.GenerateVerificationTokenResponse, error) {
	err := handler.service.GenerateVerificationToken(request.Email)
	return &pb.GenerateVerificationTokenResponse{}, err
}

func (handler *AuthenticationHandler) AccountRecovery(ctx context.Context, request *pb.AccountRecoveryRequest) (*pb.AccountRecoveryResponse, error) {
	status, err := handler.service.AccountRecovery(request.Email)
	return &pb.AccountRecoveryResponse{
		Status: status,
		Error:  err,
	}, nil
}

func (handler *AuthenticationHandler) ChangeTwoFactorAuth(ctx context.Context, request *pb.ChangeTwoFactorAuthRequest) (*pb.ChangeTwoFactorAuthResponse, error) {
	handler.service.ChangeTwoFactorAuth(request.String())
	return &pb.ChangeTwoFactorAuthResponse{}, nil
}

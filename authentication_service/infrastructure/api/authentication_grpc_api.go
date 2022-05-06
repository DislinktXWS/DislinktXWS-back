package api

import (
	"context"
	"module/authentication_service/application"
	"module/authentication_service/utils"
	pb "module/common/proto/authentication_service"
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

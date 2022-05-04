package api

import (
	"context"
	"module/authentication_service/application"
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
	auth := request.Auth
	newAuth := mapAuth(auth)
	status, err, token := handler.service.Login(newAuth)
	return &pb.LoginResponse{
		Status: status,
		Error:  err,
		Token:  token,
	}, nil
}

func (handler *AuthenticationHandler) Validate(ctx context.Context, request *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	token := request.Token
	status, err, username := handler.service.Validate(token)
	return &pb.ValidateResponse{
		Status:   status,
		Error:    err,
		Username: username,
	}, nil
}

package api

import (
	"module/authentication_service/domain"
	pb "module/common/proto/authentication_service"
)

func mapAuth(authPb *pb.Auth) *domain.Auth {
	auth := &domain.Auth{
		Username: authPb.Username,
		Password: authPb.Password,
	}
	return auth
}

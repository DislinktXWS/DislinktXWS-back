package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"module/authentication_service/domain"
	pb "module/common/proto/authentication_service"
)

func mapAuth(authPb *pb.Auth) *domain.Auth {
	id, _ := primitive.ObjectIDFromHex(authPb.Id)
	auth := &domain.Auth{
		Id:       id,
		Username: authPb.Username,
		Password: authPb.Password,
	}
	return auth
}

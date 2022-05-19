package api

import (
	"github.com/dislinktxws-back/authentication_service/domain"
	pb "github.com/dislinktxws-back/common/proto/authentication_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAuth(authPb *pb.Auth) *domain.Auth {
	id, _ := primitive.ObjectIDFromHex(authPb.Id)
	auth := &domain.Auth{
		Id:                id,
		Username:          authPb.Username,
		Password:          authPb.Password,
		VerificationToken: authPb.VerificationToken,
		Email:             authPb.Email,
	}
	return auth
}

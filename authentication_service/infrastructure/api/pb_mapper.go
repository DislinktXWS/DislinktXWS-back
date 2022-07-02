package api

import (
	"github.com/dislinktxws-back/authentication_service/domain"
	pb "github.com/dislinktxws-back/common/proto/authentication_service"
	events "github.com/dislinktxws-back/common/saga/insert_user"
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
		TwoFactorAuth:     false,
	}
	return auth
}

func mapCommandToAuth(authCommand *events.InsertUserCommand) *domain.Auth {
	id, _ := primitive.ObjectIDFromHex(authCommand.User.Id)
	auth := &domain.Auth{
		Id:                id,
		Username:          authCommand.User.Username,
		Password:          authCommand.User.Password,
		VerificationToken: authCommand.User.VerificationToken,
		Email:             authCommand.User.Email,
		TwoFactorAuth:     false,
	}
	return auth
}

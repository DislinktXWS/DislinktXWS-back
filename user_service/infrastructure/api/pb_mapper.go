package api

import (
	pb "module/common/proto/user_service"
	"module/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id: user.Id.Hex(),
	}
	return userPb
}

package api

import (
	pb "module/common/proto/user_service"
	"module/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:          user.Id.Hex(),
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		Password:    user.Password,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Email:       user.Email,
		Phone:       user.Phone,
	}
	return userPb
}

package api

import (
	"module/api_gateway/domain"
	pb "module/common/proto/user_service"
)

//ovoaj recimo mapira userRegistration na User, trebace i UserRegistration u UserAuth

func mapToUserPb(user *domain.UserRegistration) *pb.User {
	userPb := &pb.User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Email:       user.Email,
		Phone:       user.Phone,
		Biography:   user.Biography,
	}
	return userPb
}

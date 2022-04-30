package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func mapNewUser(userPb *pb.User) *domain.User {
	user := &domain.User{
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		Username:    userPb.Username,
		Password:    userPb.Password,
		DateOfBirth: userPb.DateOfBirth,
		Gender:      userPb.Gender,
		Email:       userPb.Email,
		Phone:       userPb.Phone,
	}
	return user
}

func mapEditUser(userPb *pb.User) *domain.User {
	id, _ := primitive.ObjectIDFromHex(userPb.Id)
	user := &domain.User{
		Id:          id,
		Name:        userPb.Name,
		Surname:     userPb.Surname,
		Username:    userPb.Username,
		Password:    userPb.Password,
		DateOfBirth: userPb.DateOfBirth,
		Gender:      userPb.Gender,
		Email:       userPb.Email,
		Phone:       userPb.Phone,
	}
	return user
}

func mapAddEducation(educationPb *pb.Education) *domain.Education {
	education := &domain.Education{
		School:       educationPb.School,
		Degree:       educationPb.Degree,
		FieldOfStudy: educationPb.FieldOfStudy,
		StartDate:    educationPb.StartDate,
		EndDate:      educationPb.EndDate,
		Grade:        educationPb.Grade,
		Description:  educationPb.Description,
	}
	return education
}

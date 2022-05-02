package api

import (
	"context"
	pb "module/common/proto/user_service"
	"module/user_service/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	User, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUser(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Insert(ctx context.Context, request *pb.InsertUserRequest) (*pb.InsertUserResponse, error) {
	user := mapNewUser(request.User)
	err, newUser := handler.service.Insert(user)
	if err != nil {
		return nil, err
	}
	UserPb := mapUser(newUser)
	response := &pb.InsertUserResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) EditUser(ctx context.Context, request *pb.InsertUserRequest) (*pb.EditUserResponse, error) {
	user := mapEditUser(request.User)
	_, err := handler.service.EditUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.EditUserResponse{}, nil
}

func (handler *UserHandler) AddEducation(ctx context.Context, request *pb.AddEducationRequest) (*pb.AddEducationResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	education := mapAddEducation(request.Education)
	_, err := handler.service.AddEducation(education, id)
	if err != nil {
		return nil, err
	}
	return &pb.AddEducationResponse{}, nil
}

func (handler *UserHandler) DeleteEducation(ctx context.Context, request *pb.DeleteEducationRequest) (*pb.DeleteEducationResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteEducation(id, index)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteEducationResponse{}, nil
}

func (handler *UserHandler) AddExperience(ctx context.Context, request *pb.AddExperienceRequest) (*pb.AddExperienceResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	experience := mapAddExperience(request.Experience)
	_, err := handler.service.AddExperience(experience, id)
	if err != nil {
		return nil, err
	}
	return &pb.AddExperienceResponse{}, nil
}

func (handler *UserHandler) DeleteExperience(ctx context.Context, request *pb.DeleteExperienceRequest) (*pb.DeleteExperienceResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteExperience(id, index)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteExperienceResponse{}, nil
}

func (handler *UserHandler) AddInterest(ctx context.Context, request *pb.AddInterestRequest) (*pb.AddInterestResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	interest := request.Interest
	err := handler.service.AddInterest(id, interest)
	if err != nil {
		return nil, err
	}
	return &pb.AddInterestResponse{}, nil
}

func (handler *UserHandler) DeleteInterest(ctx context.Context, request *pb.DeleteInterestRequest) (*pb.DeleteInterestResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteInterest(id, index)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteInterestResponse{}, nil
}

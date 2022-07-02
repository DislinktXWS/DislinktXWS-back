package api

import (
	"context"
	"fmt"
	pb "github.com/dislinktxws-back/common/proto/user_service"
	events "github.com/dislinktxws-back/common/saga/insert_user"
	saga "github.com/dislinktxws-back/common/saga/messaging"
	"github.com/dislinktxws-back/user_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service           *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUserHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) *UserHandler {
	o := &UserHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	o.commandSubscriber.Subscribe(o.handle)
	return o
}

func init() {
	infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(infoFile, "INFO: ", log.LstdFlags|log.Lshortfile)

	errFile, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	ErrorLogger = log.New(errFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func (handler *UserHandler) handle(command *events.InsertUserCommand) {
	reply := events.InsertUserReply{User: command.User}
	fmt.Println("USER HANDLER")
	fmt.Println(command.Type)

	switch command.Type {
	case events.RollbackInsertUser:
		fmt.Println("ROLLBACK USER INSERT")
		objectId, _ := primitive.ObjectIDFromHex(command.User.Id)
		handler.service.Delete(objectId)
		reply.Type = events.UserInsertRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	//fmt.Println("DOBAVLJANJE INFORMACIJA O SENDERU")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("Action: 1, Message: ID is not correct!")
		return nil, err
	}
	User, err := handler.service.Get(objectId)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: User not found!")
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetByUsername(ctx context.Context, request *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	username := request.Username
	User, err := handler.service.GetByUsername(username)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: User not found!")
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.GetByUsernameResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetByApiKey(ctx context.Context, request *pb.GetByApiKeyRequest) (*pb.GetByApiKeyResponse, error) {
	apiKey := request.ApiKey
	User, err := handler.service.GetByApiKey(apiKey)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: User not found!")
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.GetByApiKeyResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()

	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Users not found!")
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

func (handler *UserHandler) GetPublicUsers(ctx context.Context, request *pb.GetPublicUsersRequest) (*pb.GetPublicUsersResponse, error) {
	Users, err := handler.service.GetPublicUsers()
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Users not found!")
		return nil, err
	}
	response := &pb.GetPublicUsersResponse{
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
	users, _ := handler.service.GetAll()
	exists := false
	for _, currentUser := range users {
		if user.Id != currentUser.Id && (user.Username == currentUser.Username || user.Email == currentUser.Email) {
			exists = true
			break
		}
	}
	if !exists {
		err, newUser := handler.service.Insert(user)
		if err != nil {
			return nil, err
		}
		UserPb := mapUser(newUser)
		fmt.Println("Token:" + UserPb.VerificationToken)
		InfoLogger.Println("Action: 3, Message: User " + user.Username + " registered successfully.")
		response := &pb.InsertUserResponse{
			User: UserPb,
		}
		return response, nil
	}
	ErrorLogger.Println("Action: 4, Message: Username or email not unique!")
	return nil, nil
}

func (handler *UserHandler) EditUser(ctx context.Context, request *pb.InsertUserRequest) (*pb.EditUserResponse, error) {
	user := mapEditUser(request.User)
	_, err := handler.service.EditUser(user)
	if err != nil {
		ErrorLogger.Println("Action: 5, Message: Cannot edit user " + user.Username + "!")
		return nil, err
	}
	users, _ := handler.service.GetAll()
	exists := false
	for _, currentUser := range users {
		if user.Id != currentUser.Id && user.Username == currentUser.Username {
			exists = true
			break
		}
	}
	if !exists {
		_, err = handler.service.EditUsername(user)
		if err != nil {
			ErrorLogger.Println("Action: 5, Message: Username not unique!")
			return nil, err
		}
	}
	InfoLogger.Println("Action: 6, Message: User " + user.Username + " edited.")
	return &pb.EditUserResponse{}, nil
}

func (handler *UserHandler) SetApiKey(ctx context.Context, request *pb.SetApiKeyRequest) (*pb.SetApiKeyResponse, error) {
	apiKey, error := handler.service.SetApiKey(request.Username)
	InfoLogger.Println("Action: 7, Message: Successfully connected with agents app.")
	return &pb.SetApiKeyResponse{ApiKey: apiKey}, error
}

func (handler *UserHandler) GetEducation(ctx context.Context, request *pb.GetEducationRequest) (*pb.GetEducationResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	education, err := handler.service.GetEducation(id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Education not found!")
		return nil, err
	}

	response := &pb.GetEducationResponse{
		Education: []*pb.Education{},
	}

	for _, educ := range *education {
		e := mapEducation(&educ)
		response.Education = append(response.Education, e)
	}

	return response, nil

}

func (handler *UserHandler) AddEducation(ctx context.Context, request *pb.AddEducationRequest) (*pb.AddEducationResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	education := mapAddEducation(request.Education)
	_, err := handler.service.AddEducation(education, id)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Cannot add education!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User with id " + request.Id + " added education.")
	return &pb.AddEducationResponse{}, nil
}

func (handler *UserHandler) DeleteEducation(ctx context.Context, request *pb.DeleteEducationRequest) (*pb.DeleteEducationResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteEducation(id, index)
	if err != nil {
		ErrorLogger.Println("Action: 8, Message: Cannot delete education!")
		return nil, err
	}
	return &pb.DeleteEducationResponse{}, nil
}
func (handler *UserHandler) GetExperience(ctx context.Context, request *pb.GetExperienceRequest) (*pb.GetExperienceResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	experience, err := handler.service.GetExperience(id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: User experience not found!")
		return nil, err
	}
	response := &pb.GetExperienceResponse{
		Experience: []*pb.Experience{},
	}

	for _, exper := range *experience {
		exp := mapExperience(&exper)
		response.Experience = append(response.Experience, exp)
	}

	return response, nil
}

func (handler *UserHandler) AddExperience(ctx context.Context, request *pb.AddExperienceRequest) (*pb.AddExperienceResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	experience := mapAddExperience(request.Experience)
	_, err := handler.service.AddExperience(experience, id)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Cannot add experience!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User with id " + request.Id + " added experience.")
	return &pb.AddExperienceResponse{}, nil
}

func (handler *UserHandler) DeleteExperience(ctx context.Context, request *pb.DeleteExperienceRequest) (*pb.DeleteExperienceResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteExperience(id, index)
	if err != nil {
		ErrorLogger.Println("Action: 8, Message: Cannot delete experience!")
		return nil, err
	}
	return &pb.DeleteExperienceResponse{}, nil
}

func (handler *UserHandler) GetInterests(ctx context.Context, request *pb.GetInterestsRequest) (*pb.GetInterestsResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	interests, err := handler.service.GetInterests(id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: User interests not found!")
		return nil, err
	}
	response := &pb.GetInterestsResponse{
		Interests: interests,
	}

	return response, nil

}

func (handler *UserHandler) AddInterest(ctx context.Context, request *pb.AddInterestRequest) (*pb.AddInterestResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	interest := request.Interest
	err := handler.service.AddInterest(id, interest)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Can not add interest!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User with id " + request.Id + " added interest.")
	return &pb.AddInterestResponse{}, nil
}

func (handler *UserHandler) DeleteInterest(ctx context.Context, request *pb.DeleteInterestRequest) (*pb.DeleteInterestResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteInterest(id, index)
	if err != nil {
		ErrorLogger.Println("Action: 8, Message: Cannot delete interest!")
		return nil, err
	}
	return &pb.DeleteInterestResponse{}, nil
}

func (handler *UserHandler) GetSkills(ctx context.Context, request *pb.GetSkillsRequest) (*pb.GetSkillsResponse, error) {
	fmt.Println("USLO U GETSKILLS")
	fmt.Println("ID KORISNIKA: " + request.Id)
	id, _ := primitive.ObjectIDFromHex(request.Id)
	skills, err := handler.service.GetSkills(id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Skills not found!")
		return nil, err
	}

	response := &pb.GetSkillsResponse{
		Skills: []*pb.Skill{},
	}

	for _, skill := range *skills {
		s := mapSkill(&skill)
		response.Skills = append(response.Skills, s)
		fmt.Println("NASLO VESTINE")
		fmt.Println(s)
	}
	return response, nil

}

func (handler *UserHandler) AddSkill(ctx context.Context, request *pb.AddSkillRequest) (*pb.AddSkillResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	skill := mapAddSkill(request.Skill)
	_, err := handler.service.AddSkill(skill, id)
	if err != nil {
		ErrorLogger.Println("Action: 4, Message: Can not add skill!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User with id " + request.Id + " added skill.")
	return &pb.AddSkillResponse{}, nil
}

func (handler *UserHandler) SetPrivacy(ctx context.Context, request *pb.SetPrivacyRequest) (*pb.SetPrivacyResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	err := handler.service.SetPrivacy(request.Private, id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: User not found!")
		return nil, err
	}
	InfoLogger.Println("Action: 6, Message: User with id " + request.Id + " set privacy.")
	return &pb.SetPrivacyResponse{}, nil
}

func (handler *UserHandler) DeleteSkill(ctx context.Context, request *pb.DeleteSkillRequest) (*pb.DeleteSkillResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	index := uint(request.Index)
	err := handler.service.DeleteSkill(id, index)
	if err != nil {
		ErrorLogger.Println("Action: 8, Message: Cannot delete skill!")
		return nil, err
	}
	return &pb.DeleteSkillResponse{}, nil
}

func (handler *UserHandler) SearchProfiles(ctx context.Context, request *pb.SearchProfilesRequest) (*pb.SearchProfilesResponse, error) {
	Users, err := handler.service.SearchProfiles(request.Search)
	if err != nil {
		return nil, err
	}
	response := &pb.SearchProfilesResponse{
		Users: []*pb.User{},
	}
	for _, User := range *Users {
		current := mapUser(&User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) SetChatNotifications(ctx context.Context, request *pb.SetChatNotificationsRequest) (*pb.SetChatNotificationsResponse, error) {
	handler.service.SetChatNotifications(request.Id)
	return &pb.SetChatNotificationsResponse{}, nil
}

func (handler *UserHandler) SetPostNotifications(ctx context.Context, request *pb.SetPostNotificationsRequest) (*pb.SetPostNotificationsResponse, error) {
	handler.service.SetPostNotifications(request.Id)
	return &pb.SetPostNotificationsResponse{}, nil
}

func (handler *UserHandler) SetConnectionsNotifications(ctx context.Context, request *pb.SetConnectionsNotificationsRequest) (*pb.SetConnectionsNotificationsResponse, error) {
	handler.service.SetConnectionsNotifications(request.Id)
	return &pb.SetConnectionsNotificationsResponse{}, nil
}

func (handler *UserHandler) GetNotificationsSettings(ctx context.Context, request *pb.GetNotificationsSettingsRequest) (*pb.GetNotificationsSettingsResponse, error) {
	fmt.Println("ULAZI U GET NOTIF SETTING")
	notificationSettings, _ := handler.service.GetNotificationsSettings(request.Id)
	fmt.Println("IZASLO IZ GET NOTIF SETTING")
	fmt.Println(notificationSettings)
	return &pb.GetNotificationsSettingsResponse{
		ChatNotifications:        notificationSettings.ChatNotifications,
		ConnectionsNotifications: notificationSettings.ConnectionsNotifications,
		PostNotifications:        notificationSettings.PostNotifications,
	}, nil
}

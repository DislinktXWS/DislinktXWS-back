package api

import (
	"context"
	pb "github.com/dislinktxws-back/common/proto/post_service"
	"github.com/dislinktxws-back/post_service/application"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
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

func (handler *PostHandler) LikePost(ctx context.Context, request *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	id := request.Id
	username := request.Username
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("ID is not correct!")
		return nil, err
	}
	handler.service.LikePost(objectId, username)
	InfoLogger.Println("User " + username + " liked post with id " + id)
	return &pb.LikePostResponse{}, nil
}

func (handler *PostHandler) DislikePost(ctx context.Context, request *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	id := request.Id
	username := request.Username
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("ID is not correct!")
		return nil, err
	}
	InfoLogger.Println("User " + username + " disliked post with id " + id)
	handler.service.DislikePost(objectId, username)
	return &pb.DislikePostResponse{}, nil
}

func (handler *PostHandler) GetPostsByUser(ctx context.Context, request *pb.GetPostsByUserRequest) (*pb.GetPostsByUserResponse, error) {
	user := request.User
	Posts, err := handler.service.GetPostsByUser(user)
	if err != nil {
		return nil, err
	}
	response := &pb.GetPostsByUserResponse{
		Posts: []*pb.Post{},
	}

	for _, Post := range Posts {
		current := mapPost(Post)
		response.Posts = append(response.Posts, current)
	}

	return response, nil
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("ID is not correct!")
		return nil, err
	}
	Post, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	PostPb := mapPost(Post)
	response := &pb.GetResponse{
		Post: PostPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, Post := range Posts {
		current := mapPost(Post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) Insert(ctx context.Context, request *pb.InsertPostRequest) (*pb.InsertPostResponse, error) {
	Post := mapNewPost(request.Post)
	err := handler.service.Insert(Post)
	if err != nil {
		ErrorLogger.Println("Cannot create post!")
		return nil, err
	}
	InfoLogger.Println("User " + Post.User + " created a new post.")
	return &pb.InsertPostResponse{Id: Post.Id.String()}, nil
}

func (handler *PostHandler) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	Comment := mapNewComment(request.Comment)
	err := handler.service.CommentPost(Comment)
	if err != nil {
		ErrorLogger.Println("Cannot comment post!")
		return nil, err
	}
	InfoLogger.Println("User " + Comment.User + " commented post " + Comment.PostId)
	return &pb.CommentPostResponse{}, nil
}

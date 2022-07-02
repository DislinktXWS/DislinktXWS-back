package api

import (
	"context"
	"fmt"
	pb "github.com/dislinktxws-back/common/proto/post_service"
	"github.com/dislinktxws-back/post_service/application"
	"github.com/dislinktxws-back/post_service/tracer"
	otgo "github.com/opentracing/opentracing-go"
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
	trace       otgo.Tracer
)

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func init() {
	trace, _ = tracer.Init("post-service")
	otgo.SetGlobalTracer(trace)
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
	span := tracer.StartSpanFromContextMetadata(ctx, "LikePost")
	defer span.Finish()
	id := request.Id
	username := request.Username
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: ID is not correct!")
		log.Println("Action: 2, Message: ID is not correct!")
		return nil, err
	}
	handler.service.LikePost(objectId, username, ctx)
	InfoLogger.Println("Action: 10, Message: User " + username + " liked post with id " + id)
	log.Println("Action: 10, Message: User " + username + " liked post with id " + id)
	return &pb.LikePostResponse{}, nil
}

func (handler *PostHandler) DislikePost(ctx context.Context, request *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DislikePost")
	defer span.Finish()
	id := request.Id
	username := request.Username
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: ID is not correct!")
		log.Println("Action: 2, Message: ID is not correct!")
		return nil, err
	}
	InfoLogger.Println("Action: 10, Message: User " + username + " disliked post with id " + id)
	log.Println("Action: 10, Message: User " + username + " disliked post with id " + id)
	handler.service.DislikePost(objectId, username, ctx)
	return &pb.DislikePostResponse{}, nil
}

func (handler *PostHandler) GetPostsByUser(ctx context.Context, request *pb.GetPostsByUserRequest) (*pb.GetPostsByUserResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetPostsByUser")
	defer span.Finish()
	user := request.User
	Posts, err := handler.service.GetPostsByUser(user, ctx)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Posts not found!")
		log.Println("Action: 2, Message: Posts not found!")
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
	span := tracer.StartSpanFromContextMetadata(ctx, "Get")
	defer span.Finish()
	id := request.Id
	fmt.Println("USLO U GETPOSTBYID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ErrorLogger.Println("Action: 1, Message: ID is not correct!")
		log.Println("Action: 1, Message: ID is not correct!")
		return nil, err
	}
	Post, err := handler.service.Get(objectId, ctx)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Posts not found!")
		log.Println("Action: 2, Message: Posts not found!")
		return nil, err
	}
	PostPb := mapPost(Post)
	response := &pb.GetResponse{
		Post: PostPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAll")
	defer span.Finish()
	Posts, err := handler.service.GetAll(ctx)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Posts not found!")
		log.Println("Action: 2, Message: Posts not found!")
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
	span := tracer.StartSpanFromContextMetadata(ctx, "CreatePost")
	defer span.Finish()
	Post := mapNewPost(request.Post)
	err := handler.service.Insert(Post, ctx)
	if err != nil {
		ErrorLogger.Println("Action: 11, Message: Cannot create post!")
		log.Println("Action: 11, Message: Cannot create post!")
		return nil, err
	}
	InfoLogger.Println("Action: 12, Message: User " + Post.User + " created a new post.")
	log.Println("Action: 12, Message: User " + Post.User + " created a new post.")
	return &pb.InsertPostResponse{Id: Post.Id.String()}, nil
}

func (handler *PostHandler) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CommentPost")
	defer span.Finish()
	Comment := mapNewComment(request.Comment)
	err := handler.service.CommentPost(Comment, ctx)
	if err != nil {
		ErrorLogger.Println("Action: 13, Message: Cannot comment post!")
		log.Println("Action: 13, Message: Cannot comment post!")
		return nil, err
	}
	InfoLogger.Println("Action: 14, Message: User " + Comment.User + " commented post " + Comment.PostId)
	log.Println("Action: 14, Message: User " + Comment.User + " commented post " + Comment.PostId)
	return &pb.CommentPostResponse{}, nil
}

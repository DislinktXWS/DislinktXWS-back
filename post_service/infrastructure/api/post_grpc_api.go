package api

import (
	"context"
	pb "module/common/proto/post_service"
	"module/post_service/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (handler *PostHandler) LikePost(ctx context.Context, request *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	id := request.Id
	username := request.Username
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	handler.service.LikePost(objectId, username)
	return &pb.LikePostResponse{}, nil
}

func (handler *PostHandler) DislikePost(ctx context.Context, request *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	id := request.Id
	username := request.Username
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}
	return &pb.InsertPostResponse{Id: Post.Id.String()}, nil
}

func (handler *PostHandler) CommentPost(ctx context.Context, request *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	Comment := mapNewComment(request.Comment)
	err := handler.service.CommentPost(Comment)
	if err != nil {
		return nil, err
	}
	return &pb.CommentPostResponse{}, nil
}

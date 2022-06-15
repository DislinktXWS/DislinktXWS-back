package startup

import (
	"context"
	"fmt"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	"github.com/dislinktxws-back/post_service/infrastructure/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strings"

	//api_gw "github.com/dislinktxws-back/api_gateway/startup"
	post_service "github.com/dislinktxws-back/common/proto/post_service"
	"github.com/dislinktxws-back/post_service/application"
	"github.com/dislinktxws-back/post_service/domain"
	"github.com/dislinktxws-back/post_service/infrastructure/api"
	"github.com/dislinktxws-back/post_service/infrastructure/persistence"
	"github.com/dislinktxws-back/post_service/startup/config"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

type Response struct {
	status int64  `json:"status"`
	error  string `json:"error"`
	user   string `json:"user"`
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

const (
	QueueGroup = "post_service"
)

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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	PostStore := server.initPostStore(mongoClient)

	PostService := server.initPostService(PostStore)
	PostHandler := server.initPostHandler(PostService)

	server.startGrpcServer(PostHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPostStore(client *mongo.Client) domain.PostStore {
	store := persistence.NewPostMongoDBStore(client)
	posts, _ := store.GetAll()
	store.DeleteAll()
	for _, Post := range posts {
		err := store.Insert(Post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initPostService(store domain.PostStore) *application.PostService {
	return application.NewPostService(store)
}

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(PostHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		withServerUnaryInterceptor(),
	)
	post_service.RegisterPostServiceServer(grpcServer, PostHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// Authorization unary interceptor function to handle authorize per RPC call
func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info.FullMethod)
	if info.FullMethod != "/posts.PostService/GetAll" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		ErrorLogger.Println("Retrieving metadata failed!")
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		ErrorLogger.Println("Authorization token is not supplied!")
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]
	stringToken := strings.Split(token, "Bearer")

	// validateToken function validates the token
	authEndpoint := fmt.Sprintf("%s:%s", "authentication_service", "8000")
	authClient := service.NewAuthenticationClient(authEndpoint)
	validation, err := authClient.Validate(context.TODO(), &authentication_service.ValidateRequest{Token: strings.TrimSpace(stringToken[1])})
	if err != nil {
		log.Fatalln(err)
	}

	if validation.Status != 200 {
		ErrorLogger.Println("Cannot validate token!")
		return status.Errorf(codes.Unauthenticated, "Token is not valid!")
	}
	return nil
}

package startup

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	api_gw "module/api_gateway/startup"
	post_service "module/common/proto/post_service"
	"module/post_service/application"
	"module/post_service/domain"
	"module/post_service/infrastructure/api"
	"module/post_service/infrastructure/persistence"
	"module/post_service/startup/config"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "post_service"
)

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
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(api_gw.AuthRequired),
			grpc_recovery.StreamServerInterceptor(),
		)),
	)
	post_service.RegisterPostServiceServer(grpcServer, PostHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

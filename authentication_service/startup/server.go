package startup

import (
	"fmt"
	"github.com/dislinktxws-back/authentication_service/application"
	"github.com/dislinktxws-back/authentication_service/domain"
	"github.com/dislinktxws-back/authentication_service/infrastructure/api"
	"github.com/dislinktxws-back/authentication_service/infrastructure/persistence"
	"github.com/dislinktxws-back/authentication_service/startup/config"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
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
	QueueGroup = "authentication_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	PostStore := server.initAuthStore(mongoClient)

	PostService := server.initAuthService(PostStore)
	PostHandler := server.initAuthHandler(PostService)

	server.startGrpcServer(PostHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAuthStore(client *mongo.Client) domain.AuthenticationStore {
	store := persistence.NewAuthMongoDBStore(client)
	return store
}

func (server *Server) initAuthService(store domain.AuthenticationStore) *application.AuthenticationService {
	return application.NewAuthenticationService(store)
}

func (server *Server) initAuthHandler(service *application.AuthenticationService) *api.AuthenticationHandler {
	return api.NewAuthenticationHandler(service)
}

func (server *Server) startGrpcServer(AuthenticationHandler *api.AuthenticationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	authentication_service.RegisterAuthenticationServiceServer(grpcServer, AuthenticationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

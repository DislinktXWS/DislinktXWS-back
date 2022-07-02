package startup

import (
	"context"
	"fmt"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	message_service "github.com/dislinktxws-back/common/proto/message_service"
	"github.com/dislinktxws-back/message_service/application"
	"github.com/dislinktxws-back/message_service/domain"
	"github.com/dislinktxws-back/message_service/infrastructure/api"
	"github.com/dislinktxws-back/message_service/infrastructure/persistence"
	"github.com/dislinktxws-back/message_service/infrastructure/service"
	"github.com/dislinktxws-back/message_service/startup/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"

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
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	messageStore := server.initMessageStore(mongoClient)

	messageService := server.initMessageService(messageStore)
	messageHandler := server.initMessageHandler(messageService)

	server.startGrpcServer(messageHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.MessageDBHost, server.config.MessageDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initMessageStore(client *mongo.Client) domain.MessageStore {
	store := persistence.NewMessageMongoDBStore(client)
	/*users, _ := store.GetAll()
	store.DeleteAll()
	for _, User := range users {
		_, err := store.Insert(User)
		if err != nil {
			log.Fatal(err)
		}
	}*/
	return store
}

func (server *Server) initMessageService(store domain.MessageStore) *application.MessageService {
	return application.NewMessageService(store)
}

func (server *Server) initMessageHandler(service *application.MessageService) *api.MessageHandler {
	return api.NewMessageHandler(service)
}

func (server *Server) startGrpcServer(messageHandler *api.MessageHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
	//withServerUnaryInterceptor(),
	)
	message_service.RegisterMessageServiceServer(grpcServer, messageHandler)
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
	if info.FullMethod != "/users.UserService/GetPublicUsers" && info.FullMethod != "/users.UserService/SearchProfiles" &&
		info.FullMethod != "/users.UserService/Insert" && info.FullMethod != "/users.UserService/Get" {
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
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
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
		return status.Errorf(codes.Unauthenticated, "Token is not valid!")
	}
	return nil
}

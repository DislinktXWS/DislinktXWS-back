package startup

import (
	"context"
	"fmt"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	connection_service "github.com/dislinktxws-back/common/proto/connection_service"
	"github.com/dislinktxws-back/connection_service/application"
	"github.com/dislinktxws-back/connection_service/domain"
	"github.com/dislinktxws-back/connection_service/infrastructure/api"
	"github.com/dislinktxws-back/connection_service/infrastructure/persistence"
	"github.com/dislinktxws-back/connection_service/infrastructure/service"
	"github.com/dislinktxws-back/connection_service/startup/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"

	"github.com/neo4j/neo4j-go-driver/neo4j"
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
	QueueGroup = "connection_service"
)

func (server *Server) Start() {
	neo4jsession := server.initNeo4jSession()
	connectionStore := server.initConnectionStore(neo4jsession)

	connectionsService := server.initConnectionsService(connectionStore)
	userHandler := server.initConnectionsHandler(connectionsService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initNeo4jSession() *neo4j.Session {
	session, err := persistence.GetClient(server.config.Username, server.config.Password, server.config.Uri)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func (server *Server) initConnectionStore(client *neo4j.Session) domain.ConnectionsGraph {
	store := persistence.NewConnectionsGraph(client)
	return store
}

func (server *Server) initConnectionsService(store domain.ConnectionsGraph) *application.ConnectionsService {
	return application.NewConnectionsService(store)
}

func (server *Server) initConnectionsHandler(service *application.ConnectionsService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
	//	withServerUnaryInterceptor(),
	)
	connection_service.RegisterConnectionsServiceServer(grpcServer, connectionHandler)
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
	if info.FullMethod != "/connections.ConnectionService/InsertNewUser" {
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

package startup

import (
	"context"
	"crypto/tls"
	"fmt"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	connection_service "github.com/dislinktxws-back/common/proto/connection_service"
	saga "github.com/dislinktxws-back/common/saga/messaging"
	"github.com/dislinktxws-back/common/saga/messaging/nats"
	"github.com/dislinktxws-back/connection_service/application"
	"github.com/dislinktxws-back/connection_service/domain"
	"github.com/dislinktxws-back/connection_service/infrastructure/api"
	"github.com/dislinktxws-back/connection_service/infrastructure/persistence"
	"github.com/dislinktxws-back/connection_service/infrastructure/service"
	"github.com/dislinktxws-back/connection_service/startup/config"
	"github.com/dislinktxws-back/connection_service/tracer"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	otgo "github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"strings"
)

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init("connection-service")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
	}
}

const (
	QueueGroup = "connection_service"
)

func (server *Server) Start() {
	neo4jsession := server.initNeo4jSession()
	connectionStore := server.initConnectionStore(neo4jsession)

	commandSubscriber := server.initSubscriber(server.config.InsertUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.InsertUserReplySubject)

	connectionsService := server.initConnectionsService(connectionStore)
	userHandler := server.initConnectionsHandler(connectionsService, replyPublisher, commandSubscriber)

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

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initConnectionsService(store domain.ConnectionsGraph) *application.ConnectionsService {
	return application.NewConnectionsService(store)
}

func (server *Server) initConnectionsHandler(service *application.ConnectionsService, publisher saga.Publisher, subscriber saga.Subscriber) *api.ConnectionHandler {
	return api.NewConnectionHandler(service, publisher, subscriber)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("connectionservice.crt", "connectionservice.key")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		ErrorLogger.Println("Cannot load TLS credentials: " + err.Error())
	}

	grpcServer := grpc.NewServer(
		//grpc.Creds(tlsCredentials),
		//	withServerUnaryInterceptor(),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		)),
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

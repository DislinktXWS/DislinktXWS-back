package startup

import (
	"crypto/tls"
	"fmt"
	"github.com/dislinktxws-back/authentication_service/application"
	"github.com/dislinktxws-back/authentication_service/domain"
	"github.com/dislinktxws-back/authentication_service/infrastructure/api"
	"github.com/dislinktxws-back/authentication_service/infrastructure/persistence"
	"github.com/dislinktxws-back/authentication_service/startup/config"
	"github.com/dislinktxws-back/authentication_service/tracer"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	saga "github.com/dislinktxws-back/common/saga/messaging"
	"github.com/dislinktxws-back/common/saga/messaging/nats"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	otgo "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init("authentication-service")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
	}
}

const (
	QueueGroup = "authentication_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	PostStore := server.initAuthStore(mongoClient)

	commandSubscriber := server.initSubscriber(server.config.InsertUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.InsertUserReplySubject)

	PostService := server.initAuthService(PostStore)
	PostHandler := server.initAuthHandler(PostService, replyPublisher, commandSubscriber)

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

func (server *Server) initAuthService(store domain.AuthenticationStore) *application.AuthenticationService {
	return application.NewAuthenticationService(store)
}

func (server *Server) initInsertUserHandler(service *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateAuthCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initAuthHandler(service *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) *api.AuthenticationHandler {
	return api.NewAuthenticationHandler(service, publisher, subscriber)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("authenticationservice.crt", "authenticationservice.key")
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

func (server *Server) startGrpcServer(AuthenticationHandler *api.AuthenticationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//tlsCredentials, _ := loadTLSCredentials()

	grpcServer := grpc.NewServer(
		//grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		)),
	)

	authentication_service.RegisterAuthenticationServiceServer(grpcServer, AuthenticationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

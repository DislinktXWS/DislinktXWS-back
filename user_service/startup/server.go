package startup

import (
	"context"
	"crypto/tls"
	"fmt"
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	user_service "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/dislinktxws-back/user_service/application"
	"github.com/dislinktxws-back/user_service/domain"
	"github.com/dislinktxws-back/user_service/infrastructure/api"
	"github.com/dislinktxws-back/user_service/infrastructure/persistence"
	"github.com/dislinktxws-back/user_service/infrastructure/service"
	"github.com/dislinktxws-back/user_service/startup/config"
	"github.com/dislinktxws-back/user_service/tracer"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	otgo "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
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
	tracer, closer := tracer.Init("user-service")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
	}
}

const (
	QueueGroup = "user_service"
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
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)
	userHandler := server.initUserHandler(userService)
	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	return persistence.NewUserMongoDBStore(client)
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("userservice.crt", "userservice.key")
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

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		ErrorLogger.Println("Cannot load TLS credentials: " + err.Error())
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			serverInterceptor,
			grpc_opentracing.UnaryServerInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		)),
		//grpc.Creds(tlsCredentials),
	)

	user_service.RegisterUserServiceServer(grpcServer, userHandler)
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
		info.FullMethod != "/users.UserService/Insert" && info.FullMethod != "/users.UserService/Get" &&
		info.FullMethod != "/users.UserService/GetAll" &&
		info.FullMethod != "/users.UserService/GetByUsername" &&
		info.FullMethod != "/users.UserService/SetApiKey" &&
		info.FullMethod != "/users.UserService/GetByApiKey" {
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
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata failed!")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		ErrorLogger.Println("Action: 34, Message: Authorization token is not supplied!")
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
		ErrorLogger.Println("Action: 35, Message: Cannot validate token!")
		return status.Errorf(codes.Unauthenticated, "Token is not valid!")
	}

	return nil
}

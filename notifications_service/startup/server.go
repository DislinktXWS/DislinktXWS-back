package startup

import (
	"fmt"
	notifications_service "github.com/dislinktxws-back/common/proto/notifications_service"
	"github.com/dislinktxws-back/notification_service/application"
	"github.com/dislinktxws-back/notification_service/domain"
	"github.com/dislinktxws-back/notification_service/infrastructure/api"
	"github.com/dislinktxws-back/notification_service/infrastructure/persistence"
	"github.com/dislinktxws-back/notification_service/startup/config"
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

const (
	QueueGroup = "notifications_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	notificationsStore := server.initNotificationsStore(mongoClient)
	notificationsService := server.initNotificationsService(notificationsStore)
	notificationsHandler := server.initNotificationsHandler(notificationsService)

	server.startGrpcServer(notificationsHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.NotificationDBPort, server.config.NotificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationsStore(client *mongo.Client) domain.NotificationsStore {
	return persistence.NewNotificationsMongoDBStore(client)
}

func (server *Server) initNotificationsService(store domain.NotificationsStore) *application.NotificationsService {
	return application.NewNotificationsService(store)
}

func (server *Server) initNotificationsHandler(service *application.NotificationsService) *api.NotificationHandler {
	return api.NewNotificationsHandler(service)
}

func (server *Server) startGrpcServer(notificationsHandler *api.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	fmt.Println("POKRENUT SERVER")

	if err != nil {
		fmt.Println("Failed to listen: %v", err)
	}

	//tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		fmt.Println("Cannot load TLS credentials: " + err.Error())
	}

	grpcServer := grpc.NewServer(
	//grpc.Creds(tlsCredentials),
	)

	notifications_service.RegisterNotificationsServiceServer(grpcServer, notificationsHandler)
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("Failed to serve: %s", err)
	}
}

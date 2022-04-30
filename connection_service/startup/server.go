package startup

import (
	"fmt"
	"log"
	connection_service "module/common/proto/connection_service"
	"module/connection_service/application"
	"module/connection_service/domain"
	"module/connection_service/infrastructure/api"
	"module/connection_service/infrastructure/persistence"
	"module/connection_service/startup/config"
	"net"

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
	grpcServer := grpc.NewServer()
	connection_service.RegisterConnectionsServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

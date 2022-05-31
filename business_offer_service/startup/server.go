package startup

import (
	"fmt"
	"github.com/dislinktxws-back/business_offer_service/application"
	"github.com/dislinktxws-back/business_offer_service/domain"
	"github.com/dislinktxws-back/business_offer_service/infrastructure/api"
	"github.com/dislinktxws-back/business_offer_service/infrastructure/persistence"
	"github.com/dislinktxws-back/business_offer_service/startup/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"log"
	"net"
	business_offer_service "github.com/dislinktxws-back/common/proto/business_offer_service"
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
	QueueGroup = "business_offer_service"
)

func (server *Server) Start() {
	neo4jsession := server.initNeo4jSession()
	businessOfferStore := server.initBusinessOfferStore(neo4jsession)
	businessOfferService := server.initBusinessOfferService(businessOfferStore)
	businessOfferHandler := server.initBusinessOfferHandler(businessOfferService)

	server.startGrpcServer(businessOfferHandler)
}

func (server *Server) initNeo4jSession() *neo4j.Session {
	session, err := persistence.GetClient(server.config.Username, server.config.Password, server.config.Uri)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func (server *Server) initBusinessOfferStore(client *neo4j.Session) domain.BusinessOffersGraph {
	store := persistence.NewBusinessOffersGraph(client)
	return store
}

func (server *Server) initBusinessOfferService(store domain.BusinessOffersGraph) *application.BusinessOfferService {
	return application.NewBusinessOfferService(store)
}

func (server *Server) initBusinessOfferHandler (service *application.BusinessOfferService) *api.BusinessOfferHandler {
	return api.NewBusinessOfferHandler(service)
}

func (server *Server) startGrpcServer(businessOfferHandler *api.BusinessOfferHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	business_offer_service.RegisterBusinessOffersServiceServer(grpcServer, businessOfferHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

package startup

import (
	"log"
	"module/user_service/application"
	"module/user_service/domain"
	"module/user_service/infrastructure/api"
	"module/user_service/infrastructure/persistence"
	"module/user_service/startup/config"

	"go.mongodb.org/mongo-driver/mongo"
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
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)
	userHandler := server.initUserHandler(userService)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	users, _ := store.GetAll()
	store.DeleteAll()
	for _, User := range users {
		err := store.Insert(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewOrderService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewOrderHandler(service)
}

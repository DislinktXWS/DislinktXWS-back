package startup

import (
	"context"
	"fmt"
	"module/api_gateway/infrastructure/api"
	cfg "module/api_gateway/startup/config"
	"net/http"

	userGw "module/common/proto/user_service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	//server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	userHandler := api.NewUserHandler(userEndpoint)
	userHandler.Init(server.mux)
}

func (server *Server) Start() {

	listeningOn := server.config.Host + ":" + server.config.Port
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
	http.ListenAndServe(listeningOn, server.mux)
}

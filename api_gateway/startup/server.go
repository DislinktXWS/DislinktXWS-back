package startup

import (
	"context"
	"fmt"
	"module/api_gateway/infrastructure/api"
	cfg "module/api_gateway/startup/config"
	"net/http"

	"github.com/gorilla/handlers"

	postGw "module/common/proto/post_service"
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
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err = postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)

	//ovi handler se prave po funkcionalnosti
	registrationHandler := api.NewRegistrationHandler(userEndpoint, connectionEndpoint)
	registrationHandler.Init(server.mux)

	userConnectionsHandler := api.NewUserConnectionsHandler(userEndpoint, connectionEndpoint)
	userConnectionsHandler.Init(server.mux)

	postHandler := api.NewPostHandler(postEndpoint)
	postHandler.Init(server.mux)
}

func (server *Server) Start() {
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
	)
	listeningOn := server.config.Host + ":" + server.config.Port
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))

	http.ListenAndServe(listeningOn, ch(server.mux))

}

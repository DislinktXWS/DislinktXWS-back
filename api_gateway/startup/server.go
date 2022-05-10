package startup

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/infrastructure/api"
	cfg "github.com/dislinktxws-back/api_gateway/startup/config"
	authGw "github.com/dislinktxws-back/common/proto/authentication_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/handlers"

	connectionGw "github.com/dislinktxws-back/common/proto/connection_service"
	postGw "github.com/dislinktxws-back/common/proto/post_service"
	userGw "github.com/dislinktxws-back/common/proto/user_service"
	"net/http"
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
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainStreamInterceptor()}
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	err := authGw.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err = postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)

	if err != nil {
		panic(err)
	}

	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	err = connectionGw.RegisterConnectionsServiceHandlerFromEndpoint(context.TODO(), server.mux, connectionEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)

	registrationHandler := api.NewRegistrationHandler(userEndpoint, connectionEndpoint)
	registrationHandler.Init(server.mux)

	userConnectionsHandler := api.NewUserConnectionsHandler(userEndpoint, connectionEndpoint)
	userConnectionsHandler.Init(server.mux)

	userPostHandler := api.NewUserPostHandler(userEndpoint, postEndpoint)
	userPostHandler.Init(server.mux)

	postHandler := api.NewPostHandler(postEndpoint)
	postHandler.Init(server.mux)

	getImageHandler := api.NewGetImageHandler(postEndpoint)
	getImageHandler.Init(server.mux)

	blockedUsersHandler := api.NewBlockedUsersHandler(userEndpoint, connectionEndpoint)
	blockedUsersHandler.Init(server.mux)

	connectionReqHandler := api.NewUserConnectionRequestsHandler(userEndpoint, connectionEndpoint)
	connectionReqHandler.Init(server.mux)

	userFeedHandler := api.NewUserFeedHandler(connectionEndpoint, postEndpoint)
	userFeedHandler.Init(server.mux)
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

/*func AuthRequired(context.Context) (context.Context, error) {
	ctx := gin.Context{}
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, nil
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, nil
	}
	auth := authService.AuthenticationHandler{}
	res, err := auth.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, nil
	}

	ctx.Set("user", res.User)

	ctx.Next()
	return nil, nil
}
*/

package startup

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/infrastructure/api"
	cfg "github.com/dislinktxws-back/api_gateway/startup/config"
	authGw "github.com/dislinktxws-back/common/proto/authentication_service"
	offerGw "github.com/dislinktxws-back/common/proto/business_offer_service"
	connectionGw "github.com/dislinktxws-back/common/proto/connection_service"
	messageGw "github.com/dislinktxws-back/common/proto/message_service"
	notificationsGW "github.com/dislinktxws-back/common/proto/notifications_service"
	postGw "github.com/dislinktxws-back/common/proto/post_service"
	userGw "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/gorilla/handlers"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	//userOpts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	opts := []grpc.DialOption{grpc.WithUnaryInterceptor(
		grpc_opentracing.UnaryClientInterceptor(
			grpc_opentracing.WithTracer(otgo.GlobalTracer()),
		),
	), grpc.WithTransportCredentials(insecure.NewCredentials())}
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

	offerEndpoint := fmt.Sprintf("%s:%s", server.config.BusinessOfferHost, server.config.BusinessOfferPort)
	err = offerGw.RegisterBusinessOffersServiceHandlerFromEndpoint(context.TODO(), server.mux, offerEndpoint, opts)

	messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	err = messageGw.RegisterMessageServiceHandlerFromEndpoint(context.TODO(), server.mux, messageEndpoint, opts)
	if err != nil {
		panic(err)
	}

	notificationsEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationsOfferHost, server.config.NotificationsOfferPort)
	err = notificationsGW.RegisterNotificationsServiceHandlerFromEndpoint(context.TODO(), server.mux, notificationsEndpoint, opts)
	if err != nil {
		panic(err)
	}

}

func (server *Server) initCustomHandlers() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	businessOfferEndpoint := fmt.Sprintf("%s:%s", server.config.BusinessOfferHost, server.config.BusinessOfferPort)
	notificationsEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationsOfferHost, server.config.NotificationsOfferPort)
	authenticationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	//messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)

	registrationHandler := api.NewRegistrationHandler(userEndpoint, connectionEndpoint, authenticationEndpoint)
	registrationHandler.Init(server.mux)

	connectUserAgentHandler := api.NewConnectUserAgentsHandler(userEndpoint)
	connectUserAgentHandler.Init(server.mux)

	shareBusinessOffer := api.NewShareBusinessOfferHandler(userEndpoint, businessOfferEndpoint)
	shareBusinessOffer.Init(server.mux)

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

	userRecommendationsHandler := api.NewUserRecommendationsHandler(userEndpoint, connectionEndpoint)
	userRecommendationsHandler.Init(server.mux)

	connectionReqHandler := api.NewUserConnectionRequestsHandler(userEndpoint, connectionEndpoint)
	connectionReqHandler.Init(server.mux)

	userFeedHandler := api.NewUserFeedHandler(connectionEndpoint, postEndpoint)
	userFeedHandler.Init(server.mux)

	notificationHandler := api.NewNotificationHandler(notificationsEndpoint, connectionEndpoint, userEndpoint)
	notificationHandler.Init(server.mux)

	postNotificationHandler := api.NewPostNotificationsHandler(notificationsEndpoint, connectionEndpoint, userEndpoint)
	postNotificationHandler.Init(server.mux)

	//conversationInfoHandler := api.NewConversationInfoHandler(userEndpoint, messageEndpoint)
	//conversationInfoHandler.Init(server.mux)
}

func (server *Server) Start() {
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"https://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"}),
	)
	listeningOn := server.config.Host + ":" + server.config.Port
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
	//http.ListenAndServe(listeningOn, ch(server.mux))
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
	http.ListenAndServeTLS(listeningOn,
		"localhost.crt",
		"localhost.key",
		ch(server.mux))
}

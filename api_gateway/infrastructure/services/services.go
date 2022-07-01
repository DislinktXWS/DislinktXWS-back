package services

import (
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	business_offer "github.com/dislinktxws-back/common/proto/business_offer_service"
	connection_service "github.com/dislinktxws-back/common/proto/connection_service"
	notifications_service "github.com/dislinktxws-back/common/proto/notifications_service"
	post_service "github.com/dislinktxws-back/common/proto/post_service"
	user_service "github.com/dislinktxws-back/common/proto/user_service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(address string) user_service.UserServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return user_service.NewUserServiceClient(connection)
}

func NewAuthenticationClient(address string) authentication_service.AuthenticationServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Authentication service: %v", err)
	}
	return authentication_service.NewAuthenticationServiceClient(connection)
}

func NewPostClient(address string) post_service.PostServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return post_service.NewPostServiceClient(connection)
}

func NewConnectionClient(address string) connection_service.ConnectionsServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return connection_service.NewConnectionsServiceClient(connection)
}

func NewBusinessOfferClient(address string) business_offer.BusinessOffersServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return business_offer.NewBusinessOffersServiceClient(connection)
}

func NewNotificationClient(address string) notifications_service.NotificationsServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Notification service: %v", err)
	}
	return notifications_service.NewNotificationsServiceClient(connection)
}

func getClientConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

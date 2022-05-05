package services

import (
	"log"
	connection_service "module/common/proto/connection_service"
	post_service "module/common/proto/post_service"
	user_service "module/common/proto/user_service"

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

func NewConnectionClient(address string) connection_service.ConnectionsServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return connection_service.NewConnectionsServiceClient(connection)
}

func NewPostClient(address string) post_service.PostServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return post_service.NewPostServiceClient(connection)
}

func getClientConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

package service

import (
	authentication_service "github.com/dislinktxws-back/common/proto/authentication_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewAuthenticationClient(address string) authentication_service.AuthenticationServiceClient {
	connection, err := getClientConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return authentication_service.NewAuthenticationServiceClient(connection)
}

func getClientConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

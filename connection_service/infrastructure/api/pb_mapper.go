package api

import (
	pb "module/common/proto/connection_service"
	"module/connection_service/domain"
)

func mapNewUserConnection(userPb *pb.UserConnection) *domain.UserConnection {
	userConnection := &domain.UserConnection{
		Connected:  userPb.Connected,
		Connecting: userPb.Connecting,
	}
	return userConnection
}

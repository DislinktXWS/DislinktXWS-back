package api

import (
	pb "github.com/dislinktxws-back/common/proto/connection_service"
	"github.com/dislinktxws-back/connection_service/domain"
)

func mapNewUserConnection(userPb *pb.UserConnection) *domain.UserConnection {
	userConnection := &domain.UserConnection{
		Connected:  userPb.Connected,
		Connecting: userPb.Connecting,
	}
	return userConnection
}

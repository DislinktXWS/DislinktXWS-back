package persistence

import (
	"fmt"
	"github.com/dislinktxws-back/authentication_service/domain"
	"github.com/sec51/twofactor"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapTwoAuth(username string, totp *twofactor.Totp) *domain.TwoFactorAuth {
	id, _ := primitive.ObjectIDFromHex("")
	bytes, _ := totp.ToBytes()
	auth := &domain.TwoFactorAuth{
		Id:       id,
		Username: username,
		Totp:     bytes,
	}
	fmt.Println(auth)
	return auth
}

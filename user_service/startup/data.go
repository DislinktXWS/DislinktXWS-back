package startup

import (
	"module/user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var orders = []*domain.User{
	{
		Id:          getObjectId("623b0cc336a1d6fd8c1cf0f6"),
		Name:        "Milica",
		Surname:     "Vucinic",
		Username:    "mici",
		Password:    "mici123",
		DateOfBirth: "12.12.1999.",
		Gender:      "f",
		Email:       "mici@gmail.com",
		Phone:       "063123456",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

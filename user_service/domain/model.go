package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Username    string             `bson:"username"`
	Password    string             `bson:"password"`
	DateOfBirth string             `bson:"dateOfBirth"`
	Gender      string             `bson:"gender"`
	Email       string             `bson:"email"`
	Phone       string             `bson:"phone"`
}

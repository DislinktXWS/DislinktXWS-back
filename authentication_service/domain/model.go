package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	Username          string             `bson:"username"`
	Password          string             `bson:"password"`
	Email             string             `bson:"email"`
	VerificationToken string             `bson:"verificationToken"`
	IsVerified        bool               `bson:"isVerified"`
}

package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Auth struct {
	Id                       primitive.ObjectID `bson:"_id,omitempty"`
	Username                 string             `bson:"username"`
	Password                 string             `bson:"password"`
	Email                    string             `bson:"email"`
	VerificationToken        string             `bson:"verificationToken"`
	IsVerified               bool               `bson:"isVerified"`
	VerificationCreationTime time.Time          `bson:"verificationCreationTime"`
	TwoFactorAuth            bool               `bson:"twoFactorAuth"`
}

type TwoFactorAuth struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `bson:"username"`
	Totp     []byte             `bson:"totp"`
}

package persistence

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/authentication_service/domain"
	"github.com/dislinktxws-back/authentication_service/startup/config"
	utils "github.com/dislinktxws-back/authentication_service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

const (
	DATABASE   = "authDB"
	COLLECTION = "authentications"
)

type AuthMongoDBStore struct {
	authentications *mongo.Collection
}

func NewAuthMongoDBStore(client *mongo.Client) domain.AuthenticationStore {
	authentications := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthMongoDBStore{
		authentications: authentications,
	}
}

func (store *AuthMongoDBStore) Register(auth *domain.Auth) error {
	result, err := store.authentications.InsertOne(context.TODO(), auth)
	if err != nil {
		return err
	}
	auth.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AuthMongoDBStore) Validate(token string) (int64, string, string) {
	secretKey := config.NewConfig().JWTSecretKey
	wrapper := utils.JwtWrapper{SecretKey: secretKey, ExpirationHours: 5}
	claims, err := wrapper.ValidateToken(token)
	if err != nil {
		return http.StatusBadRequest, "Invalid token", ""
	}
	split := strings.Split(claims.Id, "\"")
	id, _ := primitive.ObjectIDFromHex(split[1])
	filter := bson.M{"_id": id}
	authentication, err := store.filterOne(filter)
	if err != nil {
		return http.StatusNotFound, "User not found", ""
	}
	return http.StatusOK, "", authentication.Id.Hex()
}

func (store *AuthMongoDBStore) Login(auth *domain.Auth) (int64, string, string) {
	filter := bson.M{"username": auth.Username}
	fmt.Println(auth.Username)
	fmt.Println(auth.Id)
	fmt.Println(auth.Password)
	authentication, err := store.filterOne(filter)
	if err != nil {
		return http.StatusNotFound, "User not found", ""
	}
	/*if auth.Password != authentication.Password {
		return http.StatusNotFound, "Wrong password", ""
	}*/
	match := utils.CheckPasswordHash(auth.Password, authentication.Password)
	if !match {
		return http.StatusNotFound, "User not found", ""
	}
	secretKey := config.NewConfig().JWTSecretKey
	wrapper := utils.JwtWrapper{SecretKey: secretKey, ExpirationHours: 5}
	token, _ := wrapper.GenerateToken(authentication)
	return http.StatusOK, "", token
}

func (store *AuthMongoDBStore) EditUsername(auth *domain.Auth) (*domain.Auth, error) {
	filter := bson.D{{}} //D je getovanje ali  po redosledu kakav je u bazi
	auths, _ := store.filter(filter)
	exists := false
	for _, currentAuth := range auths {
		if currentAuth.Id != auth.Id && currentAuth.Username == auth.Username {
			exists = true
			break
		}
	}
	if !exists {
		_, err := store.authentications.UpdateOne(
			context.TODO(),
			bson.M{"_id": auth.Id},
			bson.D{
				{"$set", bson.D{
					{"username", auth.Username},
				}},
			},
		)
		return auth, err
	}
	return auth, nil
}

func (store *AuthMongoDBStore) filter(filter interface{}) ([]*domain.Auth, error) {
	cursor, err := store.authentications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AuthMongoDBStore) filterOne(filter interface{}) (Auth *domain.Auth, err error) {
	result := store.authentications.FindOne(context.TODO(), filter)
	err = result.Decode(&Auth)
	return
}

func decode(cursor *mongo.Cursor) (authentications []*domain.Auth, err error) {
	for cursor.Next(context.TODO()) {
		var Auth domain.Auth
		err = cursor.Decode(&Auth)
		if err != nil {
			return
		}
		authentications = append(authentications, &Auth)
	}
	err = cursor.Err()
	return
}

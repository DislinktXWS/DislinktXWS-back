package persistence

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/dislinktxws-back/authentication_service/domain"
	"github.com/dislinktxws-back/authentication_service/startup/config"
	utils "github.com/dislinktxws-back/authentication_service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"net/smtp"
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

func (store *AuthMongoDBStore) GenerateVerificationToken(email string) error {
	filter := bson.M{"email": email}
	authentication, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	token := EncodeToString(6)
	store.authentications.UpdateOne(
		context.TODO(),
		bson.M{"_id": authentication.Id},
		bson.D{
			{"$set", bson.D{
				{"verificationToken", utils.HashPassword(token)},
			}},
		},
	)
	sendEmail(email, token)
	return nil
}

func sendEmail(email, token string) {
	// Sender data.
	from := "pswapoteka@gmail.com"
	password := "psw12345"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	fromMessage := fmt.Sprintf("From: <%s>\r\n", "sender@gmail.com")
	toMessage := fmt.Sprintf("To: <%s>\r\n", "recipient@gmail.com")
	subject := "Welcome to dislinkt!\r\n"
	body := "In order to login, enter this token:" + token + "\r\nWelcome!\r\n"
	msg := fromMessage + toMessage + subject + "\r\n" + body
	fmt.Println(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (store *AuthMongoDBStore) AccountRecovery(email string) (int64, string, string) {
	filter := bson.M{"email": email}
	authentication, err := store.filterOne(filter)
	if err != nil {
		return http.StatusNotFound, "User not found", ""
	}
	sendRecoveryEmail(email)
	secretKey := config.NewConfig().JWTSecretKey
	wrapper := utils.JwtWrapper{SecretKey: secretKey, ExpirationHours: 5}
	token, _ := wrapper.GenerateToken(authentication)
	return http.StatusOK, "", token
}

func sendRecoveryEmail(email string) {
	// Sender data.
	from := "pswapoteka@gmail.com"
	password := "psw12345"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	link := "http://localhost:4200/changePassword"
	subject := "Subject: Account recovery\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body>Please, follow the link where you can change your password <a href=\" " + link + "\">here </a></body></html>"
	msg := []byte(subject + mime + body)
	fmt.Println(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func (store *AuthMongoDBStore) PasswordlessLogin(verificationToken string) (int64, string, string) {
	filter := bson.D{{}}
	authentications, _ := store.filter(filter)
	for _, auth := range authentications {
		if utils.CheckPasswordHash(verificationToken, auth.VerificationToken) {
			store.authentications.UpdateOne(
				context.TODO(),
				bson.M{"_id": auth.Id},
				bson.D{
					{"$set", bson.D{
						{"isVerified", true},
					}},
				},
			)
			secretKey := config.NewConfig().JWTSecretKey
			wrapper := utils.JwtWrapper{SecretKey: secretKey, ExpirationHours: 5}
			token, _ := wrapper.GenerateToken(auth)
			return http.StatusOK, "", token
		}

	}
	return http.StatusNotFound, "User not found", ""

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
	if !authentication.IsVerified {
		return http.StatusForbidden, "User not verified", ""
	}
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

func (store *AuthMongoDBStore) ChangePassword(auth *domain.Auth) error {
	filter := bson.M{"username": auth.Username}
	authentication, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	store.authentications.UpdateOne(
		context.TODO(),
		bson.M{"_id": authentication.Id},
		bson.D{
			{"$set", bson.D{
				{"password", utils.HashPassword(auth.Password)},
			}},
		},
	)
	return nil
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

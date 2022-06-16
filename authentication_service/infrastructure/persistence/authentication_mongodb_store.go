package persistence

import (
	"context"
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dislinktxws-back/authentication_service/domain"
	"github.com/dislinktxws-back/authentication_service/startup/config"
	utils "github.com/dislinktxws-back/authentication_service/utils"
	"github.com/sec51/twofactor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

const (
	DATABASE   = "authDB"
	COLLECTION = "authentications"
)

type AuthMongoDBStore struct {
	authentications *mongo.Collection
	twoFactorAuths  *mongo.Collection
}

func NewAuthMongoDBStore(client *mongo.Client) domain.AuthenticationStore {
	authentications := client.Database(DATABASE).Collection(COLLECTION)
	twoFactorAuths := client.Database(DATABASE).Collection("twoFactorAuths")
	return &AuthMongoDBStore{
		authentications: authentications,
		twoFactorAuths:  twoFactorAuths,
	}
}

func (store *AuthMongoDBStore) VerifyTwoFactorAuthToken(username string, twoAuthToken string) (status int64, error string, JWTtoken string) {
	filter := bson.M{"username": username}
	twoFAuth, _ := store.filterOneTwoFactor(filter)
	authentication, _ := store.filterOne(filter)
	fmt.Println(twoFAuth.Username)
	fmt.Println(twoFAuth.Totp)
	otp, _ := twofactor.TOTPFromBytes(twoFAuth.Totp, username)
	err := otp.Validate(twoAuthToken)
	if err != nil {
		return http.StatusNotFound, err.Error(), ""
	}
	secretKey := config.NewConfig().JWTSecretKey
	wrapper := utils.JwtWrapper{SecretKey: secretKey, ExpirationHours: 5}
	token, _ := wrapper.GenerateToken(authentication)
	return http.StatusOK, "", token
}

func (store *AuthMongoDBStore) GetTwoFactorAuth(username string) bool {
	filter := bson.M{"username": username}
	authentication, _ := store.filterOne(filter)
	return authentication.TwoFactorAuth
}

func (store *AuthMongoDBStore) ChangeTwoFactorAuth(username string) (qrCode string, error string) {
	fmt.Println(username)
	filter := bson.M{"username": username}
	authentication, _ := store.filterOne(filter)
	if authentication.TwoFactorAuth == true {
		store.authentications.UpdateOne(
			context.TODO(),
			filter,
			bson.M{"$set": bson.M{"twoFactorAuth": false}},
		)

		store.twoFactorAuths.DeleteOne(context.TODO(), filter)
	} else {
		store.authentications.UpdateOne(
			context.TODO(),
			filter,
			bson.M{"$set": bson.M{"twoFactorAuth": true}},
		)
		auth, _ := store.filterOne(filter)
		otp, _ := twofactor.NewTOTP(auth.Email, auth.Username, crypto.SHA1, 6)
		qrBytes, _ := otp.QR()
		base64QR := base64.StdEncoding.EncodeToString(qrBytes)
		fmt.Println(otp)
		fmt.Println(*otp)

		store.InsertTwoFactAuth(mapTwoAuth(username, otp))
		res, _ := store.filterOneTwoFactor(filter)
		fmt.Println("OK")
		fmt.Println(res.Username)
		fmt.Println(res.Totp)

		fmt.Println(" inserted 2FAuth")
		return base64QR, ""
	}
	return "", ""
}

func (store *AuthMongoDBStore) InsertTwoFactAuth(twofactorAuth *domain.TwoFactorAuth) error {
	fmt.Println("INSERT METODA")
	fmt.Println(twofactorAuth)
	result, _ := store.twoFactorAuths.InsertOne(context.TODO(), twofactorAuth)
	twofactorAuth.Id = result.InsertedID.(primitive.ObjectID)
	return nil
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
				{"verificationCreationTime", time.Now()},
			}},
		},
	)
	sendEmail(email, token)
	return nil
}

func sendEmail(email, token string) {
	// Sender data.
	from := "bezbednostsomn@yahoo.com"
	password := "fcmhbptswmwtphum"

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

func (store *AuthMongoDBStore) AccountRecovery(email string) (int64, string) {
	filter := bson.M{"email": email}
	authentication, err := store.filterOne(filter)
	if err != nil {
		return http.StatusNotFound, "User not found"
	}
	token := EncodeToString(6)
	store.authentications.UpdateOne(
		context.TODO(),
		bson.M{"_id": authentication.Id},
		bson.D{
			{"$set", bson.D{
				{"verificationToken", utils.HashPassword(token)},
				{"verificationCreationTime", time.Now()},
			}},
		},
	)
	sendRecoveryEmail(email, token)
	return http.StatusOK, ""
}

func sendRecoveryEmail(email, token string) {
	// Sender data.
	from := "bezbednostsomn@yahoo.com"
	password := "fcmhbptswmwtphum"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	link := "http://localhost:4200/changePassword/" + token
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
		if utils.CheckPasswordHash(verificationToken, auth.VerificationToken) &&
			time.Now().Sub(auth.VerificationCreationTime).Minutes() <= 10 {
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

func (store *AuthMongoDBStore) Login(auth *domain.Auth) (int64, string, string, bool) {
	filter := bson.M{"username": auth.Username}
	fmt.Println(auth.Username)
	fmt.Println(auth.Id)
	fmt.Println(auth.Password)
	authentication, err := store.filterOne(filter)
	if err != nil {
		return http.StatusNotFound, "User not found", "", false
	}
	if !authentication.IsVerified {
		return http.StatusForbidden, "User not verified", "", false
	}
	match := utils.CheckPasswordHash(auth.Password, authentication.Password)
	if !match {
		return http.StatusNotFound, "User not found", "", false
	}
	secretKey := config.NewConfig().JWTSecretKey
	wrapper := utils.JwtWrapper{SecretKey: secretKey, ExpirationHours: 5}
	token, _ := wrapper.GenerateToken(authentication)
	return http.StatusOK, "", token, authentication.TwoFactorAuth
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

func (store *AuthMongoDBStore) filterOneTwoFactor(filter interface{}) (Auth *domain.TwoFactorAuth, err error) {
	result := store.twoFactorAuths.FindOne(context.TODO(), filter)
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
func decodeTwo(cursor *mongo.Cursor) (authentications []*domain.TwoFactorAuth, err error) {
	for cursor.Next(context.TODO()) {
		var Auth domain.TwoFactorAuth
		err = cursor.Decode(&Auth)
		if err != nil {
			return
		}
		authentications = append(authentications, &Auth)
	}
	err = cursor.Err()
	return
}

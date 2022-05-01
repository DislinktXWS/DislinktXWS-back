package persistence

import (
	"context"
	"fmt"
	"module/user_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "userDB"
	COLLECTION = "users"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func (store *UserMongoDBStore) AddInterest(id primitive.ObjectID, interest string) error {
	filter := bson.M{"_id": id}
	user, _ := store.filterOne(filter)
	interestsCurrent := user.Interests
	interestsCurrent = append(interestsCurrent, interest)

	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"interests", interestsCurrent}}},
		},
	)
	return err
}

func (store *UserMongoDBStore) DeleteInterest(id primitive.ObjectID, index uint) error {
	filter := bson.M{"_id": id}
	user, _ := store.filterOne(filter)
	interestsCurrent := user.Interests
	interestsCurrent = append(interestsCurrent[:index], interestsCurrent[index+1:]...)

	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"interests", interestsCurrent}}},
		},
	)
	return err
}

func (store *UserMongoDBStore) DeleteExperience(id primitive.ObjectID, index uint) error {
	filter := bson.M{"_id": id}
	user, _ := store.filterOne(filter)
	experienceCurrent := user.Experience
	experienceCurrent = append(experienceCurrent[:index], experienceCurrent[index+1:]...)

	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"experience", experienceCurrent}}},
		},
	)
	return err

}

func (store *UserMongoDBStore) AddExperience(experience *domain.Experience, id primitive.ObjectID) (*domain.Experience, error) {
	filter := bson.M{"_id": id}
	user, _ := store.filterOne(filter)
	experienceCurrent := user.Experience
	experienceCurrent = append(experienceCurrent, *experience)

	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"experience", experienceCurrent}}},
		},
	)
	return experience, err
}

func (store *UserMongoDBStore) DeleteEducation(id primitive.ObjectID, index uint) error {
	filter := bson.M{"_id": id}
	user, _ := store.filterOne(filter)
	educationCurrent := user.Education
	educationCurrent = append(educationCurrent[:index], educationCurrent[index+1:]...)

	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"education", educationCurrent}}},
		},
	)
	return err
}

func (store *UserMongoDBStore) AddEducation(education *domain.Education, id primitive.ObjectID) (*domain.Education, error) {
	filter := bson.M{"_id": id}
	user, _ := store.filterOne(filter)
	educationCurrent := user.Education
	educationCurrent = append(educationCurrent, *education)

	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"education", educationCurrent}}},
		},
	)
	return education, err
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id} //M je getovanje ali NE po redosledu kakav je u bazi
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}} //D je getovanje ali  po redosledu kakav je u bazi
	return store.filter(filter)
}

func (store *UserMongoDBStore) Insert(User *domain.User) (error, *domain.User) {

	fmt.Print("*******************USLI SMO U STORE")
	result, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err, &domain.User{}
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	//ne znam kako za ostala polja, ali skontace se kako se citav obj vraca
	return nil, User
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) EditUser(user *domain.User) (*domain.User, error) {
	_, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"name", user.Name},
				{"surname", user.Surname},
				{"username", user.Username},
				{"dateOfBirth", user.DateOfBirth},
				{"gender", user.Gender},
				{"email", user.Email},
				{"phone", user.Phone}}},
		},
	)
	return user, err
}

func decode(cursor *mongo.Cursor) (orders []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var Order domain.User
		err = cursor.Decode(&Order)
		if err != nil {
			return
		}
		orders = append(orders, &Order)
	}
	err = cursor.Err()
	return
}

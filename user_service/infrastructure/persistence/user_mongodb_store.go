package persistence

import (
	"context"
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

func (store *UserMongoDBStore) Insert(User *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
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

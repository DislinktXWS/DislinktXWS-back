package persistence

import (
	"context"
	"module/post_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "postDB"
	COLLECTION = "posts"
)

type PostMongoDBStore struct {
	posts *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDBStore{
		posts: posts,
	}
}

func (store *PostMongoDBStore) LikePost(id primitive.ObjectID, username string) {
	filter := bson.M{"_id": id}
	removed := false
	post, _ := store.filterOne(filter)
	currentDislikes := post.Dislikes
	for i, users := range currentDislikes {
		if username == users {
			currentDislikes = append(currentDislikes[:i], currentDislikes[i+1:]...)
		}
	}
	currentLikes := post.Likes
	for i, users := range currentLikes {
		if username == users {
			currentLikes = append(currentLikes[:i], currentLikes[i+1:]...)
			removed = true
		}
	}
	if !removed {
		currentLikes = append(currentLikes, username)
	}
	store.posts.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", bson.D{{"dislikes", currentDislikes}}}})
	store.posts.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", bson.D{{"likes", currentLikes}}}})
}

func (store *PostMongoDBStore) DislikePost(id primitive.ObjectID, username string) {
	filter := bson.M{"_id": id}
	removed := false
	post, _ := store.filterOne(filter)
	currentLikes := post.Likes
	for i, users := range currentLikes {
		if username == users {
			currentLikes = append(currentLikes[:i], currentLikes[i+1:]...)
		}
	}
	currentDislikes := post.Dislikes
	for i, users := range currentDislikes {
		if username == users {
			currentDislikes = append(currentDislikes[:i], currentDislikes[i+1:]...)
			removed = true
		}
	}
	if !removed {
		currentDislikes = append(currentDislikes, username)
	}
	store.posts.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", bson.D{{"dislikes", currentDislikes}}}})
	store.posts.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", bson.D{{"likes", currentLikes}}}})
}

func (store *PostMongoDBStore) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id} //M je getovanje ali NE po redosledu kakav je u bazi
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}} //D je getovanje ali  po redosledu kakav je u bazi
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetPostsByUser(user string) ([]*domain.Post, error) {
	filter := bson.M{"user": user}
	return store.filter(filter)
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (Post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&Post)
	return
}

func (store *PostMongoDBStore) Insert(Post *domain.Post) error {
	result, err := store.posts.InsertOne(context.TODO(), Post)
	if err != nil {
		return err
	}
	Post.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) DeleteAll() {
	store.posts.DeleteMany(context.TODO(), bson.D{{}})
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var Post domain.Post
		err = cursor.Decode(&Post)
		if err != nil {
			return
		}
		posts = append(posts, &Post)
	}
	err = cursor.Err()
	return
}

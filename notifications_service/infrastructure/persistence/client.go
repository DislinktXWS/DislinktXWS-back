package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	fmt.Println("MONGO URI")
	fmt.Println(uri)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

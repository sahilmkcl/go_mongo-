package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID ` bson:"_id,omitempty"`
	Name     string             `"json:"name"`
	LastName string             `"json":"lastName"`
}

func main() {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal(err)
	}
	defer close(client, context, cancel)
	collection := client.Database("mongodatabase").Collection("users")
	// u := User{
	// 	Name:     "sahil",
	// 	LastName: "shaikh",
	// }
	// collection.InsertOne(context, u)
	var result User
	err = collection.FindOne(context, bson.D{primitive.E{Key: "name", Value: "sahil"}}).Decode(&result)
	fmt.Println(result)
}

func createConnection(url string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx, cancel, err

}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

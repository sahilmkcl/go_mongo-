package database

import (
	"Go_server/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func CreateUser(user model.User) {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	checkError(err)
	defer close(client, context, cancel)
	collection := client.Database("mongodatabase").Collection("users")
	_, err = collection.InsertOne(context, user)
	checkError(err)
}

func checkError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}

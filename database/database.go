package database

import (
	"Go_server/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = "omkar"
var col = "users"

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
	CheckError(err)
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(col)
	_, err = collection.InsertOne(context, user)
	CheckError(err)
}

func CheckError(er error) {
	if er != nil {
		log.Fatal(er)
	}
}

func GetUser() []model.User {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	CheckError(err)
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(col)
	res, err := collection.Find(context, bson.D{})
	CheckError(err)
	var users []model.User
	if err = res.All(context, &users); err != nil {
		log.Fatal(err)
	}
	return users
}

func FindUser(name string) (model.User, error) {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	CheckError(err)
	defer close(client, context, cancel)
	collection := client.Database(db).Collection(col)
	var user model.User
	err = collection.FindOne(context, bson.M{"name": name}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func UpdateUser(update model.Update) error {
	client, context, cancel, err := createConnection("mongodb://localhost:27017/")
	CheckError(err)
	defer close(client, context, cancel)
	log.Println(update)
	collection := client.Database(db).Collection(col)
	_, err = collection.UpdateOne(context, bson.M{update.ToUpdate: update.OldValue}, bson.D{{"$set",
		bson.D{primitive.E{Key: update.ToUpdate, Value: update.NewValue}}}},
	)
	if err != nil {
		return err
	} else {
		return nil
	}
}

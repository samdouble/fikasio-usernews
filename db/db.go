package db

import (
    "context"
    "log"
    "os"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    bson "go.mongodb.org/mongo-driver/bson"
)

var client *mongo.Client
var ctx context.Context

type User struct {
    EmailAddress string `json: emailAddress,omitempty bson:"emailAddress,omitempty"`
    Language string `json: language,omitempty bson:"language,omitempty"`
    Name string `json: name,omitempty bson:"name,omitempty"`
    UnsubscribedFromDailyNews bool `json: unsubscribedFromDailyNews,omitempty bson:"unsubscribedFromDailyNews,omitempty"`
}

func Connect() *mongo.Client {
	if client == nil {
        mongodbURL := os.Getenv("MONGODB_URL")
        var connectionError error
        ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
        client, connectionError = mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL))
        if connectionError != nil {
            log.Fatal(connectionError)
        }
	}
    return client
}

func Disconnect() {
    client.Disconnect(ctx)
}

func GetUsers() []User {
    client := Connect()
    database := client.Database("app")
    collection := database.Collection("users")
    var users []User
    cur, err := collection.Find(ctx, bson.D{{"unsubscribedFromDailyNews", bson.D{{"$ne", true}}}})
    if err != nil {
        log.Fatal(err)
    }
    cur.All(ctx, &users)
    cur.Close(ctx)
    return users
}

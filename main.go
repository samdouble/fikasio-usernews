package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/samdouble/greetings"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    bson "go.mongodb.org/mongo-driver/bson"
)

type MyEvent struct {
    Name string `json:"name"`
}

type User struct {
    EmailAddress string `json: emailAddress,omitempty bson:"emailAddress,omitempty"`
    Language string `json: language,omitempty bson:"language,omitempty"`
    Name string `json: name,omitempty bson:"name,omitempty"`
    UnsubscribedFromDailyNews bool `json: unsubscribedFromDailyNews,omitempty bson:"unsubscribedFromDailyNews,omitempty"`
}

var client *mongo.Client
func ConnectToMongoDB() {
	if client != nil {
		return
	}
	var connectionError error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, connectionError = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
    if connectionError != nil {
        log.Fatal(connectionError)
    }
    defer client.Disconnect(ctx)
    database := client.Database("app")
    collection := database.Collection("users")
    var users []User
    cur, err := collection.Find(ctx, bson.D{{"unsubscribedFromDailyNews", bson.D{{"$ne", true}}}})
    cur.All(ctx, &users)
    cur.Close(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(users)
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    names := []string{ "Gladys", "Samantha", "Darrin" }
    message, _ := greetings.Hellos(names)
    fmt.Println(message)
    ConnectToMongoDB()
    return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
    lambda.Start(HandleRequest)
}

package db

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    bson "go.mongodb.org/mongo-driver/bson"
)

var client *mongo.Client
func ConnectToMongoDB() *mongo.Client {
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
    return client
}

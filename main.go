package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/samdouble/db"
    "github.com/samdouble/email"
    "github.com/samdouble/greetings"
)

type MyEvent struct {
    Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    names := []string{"Gladys", "Samantha", "Darrin"}
    message, _ := greetings.Hellos(names)
    fmt.Println(message)
    users := db.GetUsers()
    fmt.Println(users)
    email.Send()
    return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
    lambda.Start(HandleRequest)
    defer db.Disconnect()
}

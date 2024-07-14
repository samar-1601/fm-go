package main

// architecture, the flow goes like trigger -> API Gateway -> Lambda functions -> response -> s3 -> CDN -> endUser
// lambda functions are event based executable functions which are used in serverless deployment

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// takes a payload and does something
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}
	return fmt.Sprintf("Successfully called by = %s", event.Username), nil
}

func main() {
	myApp := app.NewApp()
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}

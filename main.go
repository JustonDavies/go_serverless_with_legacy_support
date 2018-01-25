package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrParameterNotProvided = errors.New("The required parameters were not provided in the HTTP body.")
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n\tBody:%s", request.RequestContext.RequestID, request.Body)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrParameterNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "AWS Lambda with Go is pretty great, don't you think " + request.Body + "?",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
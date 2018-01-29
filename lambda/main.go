package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrInvalidInput is thrown when the request body does not describe a person
	ErrInvalidInput = errors.New("Provided HTTP body isn't a valid Person JSON object")
)

// Person describes a person (duh...)
type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if len(request.RequestContext.RequestID) > 0 {
		log.Printf("Processing lambda request with id %s\n", request.RequestContext.RequestID)
	}

	var person Person
	err := json.Unmarshal([]byte(request.Body), &person)

	if err != nil {
		return events.APIGatewayProxyResponse{}, ErrInvalidInput
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + person.Name,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

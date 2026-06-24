package handler

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func respond(code int, body any) (events.APIGatewayProxyResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal response body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	headers := make(map[string]string, 2)
	headers["Content-Type"] = "application/json"
	headers["Access-Control-Allow-Origin"] = "*" // TODO: tighten with auth

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    headers,
		Body:       string(jsonBody),
	}, nil
}

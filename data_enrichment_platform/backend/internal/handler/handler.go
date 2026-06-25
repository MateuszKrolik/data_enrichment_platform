package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	// Context needs to be first for lambda SAM
	Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

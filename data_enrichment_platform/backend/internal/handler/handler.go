package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	Handle(req events.APIGatewayProxyRequest, ctx context.Context) (events.APIGatewayProxyResponse, error)
}

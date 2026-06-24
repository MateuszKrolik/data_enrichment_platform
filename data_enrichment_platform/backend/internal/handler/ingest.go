package handler

import (
	"context"
	dtos "data_enrichment_platform/internal/dto"
	"data_enrichment_platform/internal/queue"
	"data_enrichment_platform/internal/store"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type IngestHandler struct {
	Store store.Store
	Queue queue.Queue
}

func NewIngestHandler(s store.Store, q queue.Queue) Handler {
	return &IngestHandler{
		Store: s,
		Queue: q,
	}
}

func (i *IngestHandler) Handle(req events.APIGatewayProxyRequest, ctx context.Context) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case http.MethodPost:
		return i.handleSubmitJob(req, ctx)
	case http.MethodGet:
		return i.handleGetJob(req, ctx)
	default:
		message := fmt.Sprintf("Method %s not allowed!", req.HTTPMethod)
		return respond(http.StatusMethodNotAllowed, dtos.ErrorResponse{Message: message})
	}
}

func (i *IngestHandler) handleGetJob(req events.APIGatewayProxyRequest, ctx context.Context) (events.APIGatewayProxyResponse, error) {
	jobID := req.PathParameters["id"]
	if jobID == "" {
		return respond(http.StatusBadRequest, dtos.ErrorResponse{Message: "Path parameter 'id' cannot be empty!"})
	}

	job, err := i.Store.GetJob(jobID, ctx)
	if err != nil {
		return respond(http.StatusInternalServerError, dtos.ErrorResponse{Message: err.Error()})
	}

	return respond(http.StatusAccepted, job)
}

func (i *IngestHandler) handleSubmitJob(req events.APIGatewayProxyRequest, ctx context.Context) (events.APIGatewayProxyResponse, error) {
	panic("unimplemented")
}

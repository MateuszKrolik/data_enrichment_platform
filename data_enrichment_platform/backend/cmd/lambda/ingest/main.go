package main

import (
	"context"
	"data_enrichment_platform/internal/handler"
	"data_enrichment_platform/internal/queue"
	"data_enrichment_platform/internal/store"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	ctx := context.Background()

	isProd, err := strconv.ParseBool(mustEnv("IS_PROD"))
	if err != nil {
		log.Fatalf("Invalid 'IS_PROD' environment variable!")
	}

	dynamoClient := store.NewDynamoClient(ctx, isProd, mustEnv("DYNAMO_BASE_ENDPOINT"))
	sqsClient := queue.NewSQSClient(ctx)

	db := store.NewDynamoStore(dynamoClient, mustEnv("TABLE_NAME"))
	q := queue.NewSQSQueue(sqsClient, mustEnv("QUEUE_URL"))

	h := handler.NewIngestHandler(db, q)

	lambda.Start(h.Handle)
}

func mustEnv(key string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return envVar
}

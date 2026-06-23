package main

import (
	"context"
	"data_enrichment_platform/internal/store"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	ctx := context.Background()

	dynamoClient := store.NewDynamoClient(ctx)
	db := store.NewDynamoStore(dynamoClient, mustEnv("TABLE_NAME"))

	// TODO: Wire up handler
	_ = db
	lambda.Start(func(ctx context.Context) error { return nil })
}

func mustEnv(key string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return envVar
}

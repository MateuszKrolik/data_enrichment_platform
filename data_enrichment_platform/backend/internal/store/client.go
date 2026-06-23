package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoClient(ctx context.Context) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Unable to load AWS config: %v", err)
	}
	return dynamodb.NewFromConfig(cfg)
}

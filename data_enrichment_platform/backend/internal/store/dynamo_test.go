package store

import (
	"context"
	"data_enrichment_platform/internal/models"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load()
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Loading default config failed: %v", err)
	}
	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:8000")
	})
	store := NewDynamoStore(client, mustEnv("TABLE_NAME"))
	err = store.CreateJob(&models.Job{
		JobID:     "1",
		Status:    models.Completed,
		Input:     map[string]string{},
		Results:   map[string]any{},
		Error:     "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		TTL:       time.Now().Add(24 * time.Hour).Unix(),
	}, ctx)
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}
}

func mustEnv(key string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return envVar
}

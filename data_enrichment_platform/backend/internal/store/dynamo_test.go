package store

import (
	"context"
	"data_enrichment_platform/internal/enum"
	"data_enrichment_platform/internal/model"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var jobID string = uuid.MustParse("f6210e35-884c-497f-9809-589635fef083").String()

func TestCreate(t *testing.T) {
	//GIVEN
	store, ctx := setupTest(t)
	job := &model.Job{
		JobID:  jobID,
		Status: enum.Pending,
		Input:  make(map[string]string, 1),
	}

	//WHEN
	err := store.CreateJob(job, ctx)

	//THEN
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}
}

func TestGet(t *testing.T) {
	//GIVEN
	store, ctx := setupTest(t)

	//WHEN
	_, err := store.GetJob(jobID, ctx)

	//THEN
	if err != nil {
		log.Fatalf("Error getting job: %v", err)
	}
}

func setupTest(t *testing.T) (Store, context.Context) {
	t.Helper()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Loading environment variables failed: %v", err)
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Loading default AWS config failed: %v", err)
	}
	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(mustEnv("DYNAMO_BASE_ENDPOINT"))
	})
	return NewDynamoStore(client, mustEnv("TABLE_NAME")), ctx
}

func mustEnv(key string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return envVar
}

package store

import (
	"context"
	"data_enrichment_platform/internal/models"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoStore struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoStore(client *dynamodb.Client, tableName string) Store {
	return &DynamoStore{
		client,
		tableName,
	}
}

func (d *DynamoStore) CreateJob(job *models.Job, ctx context.Context) error {
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	job.TTL = time.Now().Add(24 * time.Hour).Unix()

	item, err := attributevalue.MarshalMap(job)
	if err != nil {
		return fmt.Errorf("Failed marshal operation for create job: %w", err)
	}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(d.tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(jobId)"),
	})

	if err != nil {
		return fmt.Errorf("Failed put item operation for create job: %w", err)
	}

	return nil
}

func (d *DynamoStore) GetJob(jobId string, ctx context.Context) (*models.Job, error) {
	panic("unimplemented")
}

func (d *DynamoStore) UpdateJob(job *models.Job, ctx context.Context) error {
	panic("unimplemented")
}

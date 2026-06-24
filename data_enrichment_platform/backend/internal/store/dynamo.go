package store

import (
	"context"
	"data_enrichment_platform/internal/model"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (d *DynamoStore) CreateJob(job *model.Job, ctx context.Context) error {
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

func (d *DynamoStore) GetJob(jobId string, ctx context.Context) (*model.Job, error) {
	key := make(map[string]types.AttributeValue, 1)
	key["jobId"] = &types.AttributeValueMemberS{Value: jobId}

	result, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key:       key,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to get job: %w", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("Job not found: %s", jobId)
	}

	var job model.Job
	if err := attributevalue.UnmarshalMap(result.Item, &job); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal job: %w", err)
	}

	return &job, nil
}

func (d *DynamoStore) UpdateJob(job *model.Job, ctx context.Context) error {
	panic("unimplemented")
}

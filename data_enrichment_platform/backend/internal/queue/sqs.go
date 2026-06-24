package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSQueue struct {
	client   *sqs.Client
	queueUrl string
}

func NewSQSQueue(client *sqs.Client, queueUrl string) Queue {
	return &SQSQueue{
		client:   client,
		queueUrl: queueUrl,
	}
}

type jobMessage struct {
	JobID string `json:"jobId"`
}

func (s *SQSQueue) SendJob(jobID string, ctx context.Context) error {
	body, err := json.Marshal(jobMessage{JobID: jobID})
	if err != nil {
		return fmt.Errorf("Error while marshaling job message: %w", err)
	}
	_, err = s.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.queueUrl),
		MessageBody: aws.String(string(body)),
	})
	if err != nil {
		return fmt.Errorf("Error while sending message to SQS: %w", err)
	}
	return nil
}

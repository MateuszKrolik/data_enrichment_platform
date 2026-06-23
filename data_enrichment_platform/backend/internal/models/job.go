package models

import "time"

type Job struct {
	JobID     string            `dynamodbav:"jobId"`
	Status    JobStatus         `dynamodbav:"status"`
	Input     map[string]string `dynamodbav:"input"`
	Results   map[string]any    `dynamodbav:"results,omitempty"`
	Error     string            `dynamodbav:"error,omitempty"`
	CreatedAt time.Time         `dynamodbav:"createdAt"`
	UpdatedAt time.Time         `dynamodbav:"updatedAt"`
	TTL       int64             `dynamodbav:"ttl"`
}

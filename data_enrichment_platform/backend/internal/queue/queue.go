package queue

import "context"

type Queue interface {
	SendJob(jobID string, ctx context.Context) error
}

package store

import (
	"context"
	"data_enrichment_platform/internal/model"
)

type Store interface {
	CreateJob(job *model.Job, ctx context.Context) error
	GetJob(jobId string, ctx context.Context) (*model.Job, error)
	UpdateJob(job *model.Job, ctx context.Context) error
}

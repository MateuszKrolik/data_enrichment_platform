package store

import (
	"context"
	"data_enrichment_platform/internal/models"
)

type Store interface {
	CreateJob(job *models.Job, ctx context.Context) error
	GetJob(jobId string, ctx context.Context) (*models.Job, error)
	UpdateJob(job *models.Job, ctx context.Context) error
}

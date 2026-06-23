package models

type JobStatus string

const (
	Pending   JobStatus = "PENDING"
	Running   JobStatus = "RUNNING"
	Completed JobStatus = "COMPLETED"
	Failed    JobStatus = "FAILED"
)

func (s JobStatus) IsValid() bool {
	switch s {
	case Pending, Running, Completed, Failed:
		return true
	default:
		return false
	}
}

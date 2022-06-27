package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"time"
)

// Pipeline Pipeline operations
type Pipeline interface {
	GetByProcessId(processId string) v1.Pipeline
	GetProcessByCompanyIdAndProcessId(companyId, processId string) v1.Process
	GetStatusCount(companyId string, fromDate, toDate time.Time) v1.PipelineStatusCount
}

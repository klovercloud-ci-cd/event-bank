package repository

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"time"
)

// ProcessEventRepository Process event repository operations.
type ProcessEventRepository interface {
	Store(data v1.PipelineProcessEvent)
	GetByCompanyId(companyId string) map[string]interface{}
	GetByCompanyIdAndProcessId(companyId, processId string, option v1.ProcessQueryOption) ([]v1.PipelineProcessEvent, int64)
	DequeueByCompanyIdAndUserId(companyId, userId string) map[string]interface{}
	DequeueByCompanyIdAndUserIdAndTime(companyId, userId string, from time.Time) map[string]interface{}
}

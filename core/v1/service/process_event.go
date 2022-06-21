package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"time"
)

// ProcessEvent Process Event operations.
type ProcessEvent interface {
	Store(data v1.PipelineProcessEvent)
	GetByCompanyId(companyId string) map[string]interface{}
	GetByCompanyIdAndProcessId(companyId, processId string, option v1.ProcessQueryOption) ([]v1.PipelineProcessEvent, int64)
	DequeueByCompanyIdAndUserId(companyId, userId string) map[string]interface{}
	DequeueByCompanyIdAndUserIdAndTime(companyId, userId string, createdTime time.Time) map[string]interface{}
	ReadEventByCompanyIdAndUserId(c chan map[string]interface{}, companyId, userId string)
	ReadEventByCompanyIdAndUserIdAndTime(c chan map[string]interface{}, companyId, userId string, from time.Time)
}

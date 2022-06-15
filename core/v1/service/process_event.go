package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
)

// ProcessEvent Process Event operations.
type ProcessEvent interface {
	Store(data v1.PipelineProcessEvent)
	GetByCompanyId(companyId string) map[string]interface{}
	GetByCompanyIdAndProcessId(companyId, processId string, option v1.ProcessQueryOption) []v1.PipelineProcessEvent
	DequeueByCompanyId(companyId string) map[string]interface{}
	ReadEventByCompanyId(c chan map[string]interface{}, companyId string)
}

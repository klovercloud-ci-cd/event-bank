package repository

import v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"

// ProcessEventRepository Process event repository operations.
type ProcessEventRepository interface {
	Store(data v1.PipelineProcessEvent)
	GetByCompanyId(companyId string) map[string]interface{}
	DequeueByCompanyId(companyId string) map[string]interface{}
}

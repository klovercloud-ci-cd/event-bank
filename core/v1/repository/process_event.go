package repository

import v1 "github.com/klovercloud-ci/core/v1"

// ProcessEventRepository Process event repository operations.
type ProcessEventRepository interface {
	Store(data v1.PipelineProcessEvent)
	GetByProcessId(processId string) map[string]interface{}
	DequeueByProcessId(processId string) map[string]interface{}
}

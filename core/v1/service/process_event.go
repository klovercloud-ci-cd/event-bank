package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
)

// ProcessEvent Process Event operations.
type ProcessEvent interface {
	Store(data v1.PipelineProcessEvent)
	GetByProcessId(processId string) map[string]interface{}
	DequeueByProcessId(processId string) map[string]interface{}
	ReadEventByProcessId(c chan map[string]interface{}, processId string)
}

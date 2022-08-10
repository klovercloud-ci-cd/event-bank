package repository

import v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"

// ProcessFootmarkRepository Process log footmark Repository operations.
type ProcessFootmarkRepository interface {
	Store(processFootmark v1.ProcessFootmark)
	GetByProcessId(processId string) []v1.ProcessFootmark
	GetByProcessIdAndStep(processId, step string) []v1.ProcessFootmark
	GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark string, claim int) *v1.ProcessFootmark
}

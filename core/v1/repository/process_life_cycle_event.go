package repository

import v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"

// ProcessLifeCycleEventRepository Process life cycle event repository operations.
type ProcessLifeCycleEventRepository interface {
	Store(data []v1.ProcessLifeCycleEvent)
	Get(count int64) []v1.ProcessLifeCycleEvent
	PullPausedAndAutoTriggerEnabledResourcesByAgentName(count int64, agent string) []v1.ProcessLifeCycleEvent
	PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count int64, stepType string) []v1.ProcessLifeCycleEvent
	GetByProcessId(processId string) []v1.ProcessLifeCycleEvent
	UpdateClaim(processId,step,status string) error
}

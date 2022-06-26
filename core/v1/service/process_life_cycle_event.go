package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"time"
)

// ProcessLifeCycleEvent Process Life Cycle Event operations.
type ProcessLifeCycleEvent interface {
	Store(events []v1.ProcessLifeCycleEvent)
	GetByProcessIdAndStep(processId, step string) v1.ProcessLifeCycleEvent
	PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count int64, stepType string) []v1.ProcessLifeCycleEvent
	PullPausedAndAutoTriggerEnabledResourcesByAgentName(count int64, agent string) []v1.DeployableResource
	GetByProcessId(processId string) []v1.ProcessLifeCycleEvent
	UpdateClaim(companyId, processId, step, status string) error
	UpdateStatusesByTime(time time.Time)
}

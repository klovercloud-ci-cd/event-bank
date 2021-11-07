package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
)

type ProcessLifeCycleEvent interface {
	Store( events []v1.ProcessLifeCycleEvent)
	PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count int64,stepType string)[]v1.ProcessLifeCycleEvent
	PullPausedAndAutoTriggerEnabledResourcesByAgentName(count int64,agent string)[]v1.DeployableResource
}




package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
)

var eventsData []v1.ProcessLifeCycleEvent

// InitProcessLifeCycleEventData Return mock process life cycle event list
func InitProcessLifeCycleEventData() []v1.ProcessLifeCycleEvent {
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "24033301-5186-4de3-8ab2-469b8d717e45",
		Step:      "build",
		StepType:  "BUILD",
		Status:    "completed",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "24033301-5186-4de3-8ab2-469b8d717e45",
		Step:      "interstep",
		StepType:  "INTERMEDIARY",
		Status:    "completed",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "24033301-5186-4de3-8ab2-469b8d717e45",
		Step:      "deployDev",
		StepType:  "DEPLOY",
		Status:    "completed",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "24033301-5186-4de3-8ab2-469b8d717e45",
		Step:      "jenkinsjob",
		StepType:  "JENKINS_JOB",
		Status:    "non_initialized",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "12c9cdcd-55e0-4b21-a124-9329eabe991a",
		Step:      "build",
		StepType:  "BUILD",
		Status:    "completed",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "12c9cdcd-55e0-4b21-a124-9329eabe991a",
		Step:      "interstep",
		StepType:  "INTERMEDIARY",
		Status:    "completed",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "12c9cdcd-55e0-4b21-a124-9329eabe991a",
		Step:      "deployDev",
		StepType:  "DEPLOY",
		Status:    "completed",
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "12c9cdcd-55e0-4b21-a124-9329eabe991a",
		Step:      "jenkinsjob",
		StepType:  "JENKINS_JOB",
		Status:    "non_initialized",
	})
	return eventsData
}

// NewMockProcessLifeCycleEventRepository returns ProcessLifeCycleEventRepository type object
func NewMockProcessLifeCycleEventRepository() repository.ProcessLifeCycleEventRepository {
	manager := GetMockDmManager()
	manager.Db.Drop(context.Background())
	return &processLifeCycleRepository{
		manager: GetMockDmManager(),
		timeout: 3000,
	}
}
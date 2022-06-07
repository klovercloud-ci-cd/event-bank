package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"time"
)

var eventsData []v1.ProcessLifeCycleEvent

// InitProcessLifeCycleEventData Return mock process life cycle event list
func InitProcessLifeCycleEventData() []v1.ProcessLifeCycleEvent {
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "1",
		Step:      "build",
		StepType:  "BUILD",
		Status:    "active",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "2",
		Step:      "interstep",
		StepType:  "INTERMEDIARY",
		Status:    "active",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "3",
		Step:      "deployDev",
		StepType:  "DEPLOY",
		Status:    "completed",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "4",
		Step:      "jenkinsjob",
		StepType:  "JENKINS_JOB",
		Status:    "non_initialized",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "5",
		Step:      "build",
		StepType:  "BUILD",
		Status:    "completed",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "6",
		Step:      "interstep",
		StepType:  "INTERMEDIARY",
		Status:    "completed",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "7",
		Step:      "deployDev",
		StepType:  "DEPLOY",
		Status:    "completed",
		CreatedAt: time.Now().UTC(),
	})
	eventsData = append(eventsData, v1.ProcessLifeCycleEvent{
		ProcessId: "8",
		Step:      "jenkinsjob",
		StepType:  "JENKINS_JOB",
		Status:    "non_initialized",
		CreatedAt: time.Now().UTC(),
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

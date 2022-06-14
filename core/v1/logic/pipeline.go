package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"time"
)

type pipelineService struct {
	processLifeCycleEventRepository repository.ProcessLifeCycleEventRepository
}

func (p pipelineService) GetStatusCount(companyId string, fromDate, toDate time.Time) v1.PipelineStatusCount {
	events := p.processLifeCycleEventRepository.GetByCompanyId(companyId, fromDate, toDate)
	processStatusMap := GetProcessStatusMapFromEvents(events)
	var completed, failed, running, paused, nonInitialized int64
	for _, each := range processStatusMap {
		if each == enums.COMPLETED {
			completed++
		} else if each == enums.FAILED {
			failed++
		} else if each == enums.ACTIVE {
			running++
		} else if each == enums.PAUSED {
			paused++
		} else if each == enums.NON_INITIALIZED {
			nonInitialized++
		}
	}
	return v1.PipelineStatusCount{
		Pipeline: struct {
			Completed      int64 `json:"completed"`
			Failed         int64 `json:"failed"`
			Running        int64 `json:"running"`
			Paused         int64 `json:"paused"`
			NonInitialized int64 `json:"nonInitialized"`
		}(struct {
			Completed      int64
			Failed         int64
			Running        int64
			Paused         int64
			NonInitialized int64
		}{Completed: completed, Failed: failed, Running: running, Paused: paused, NonInitialized: nonInitialized}),
	}
}

func GetProcessStatusMapFromEvents(events []v1.ProcessLifeCycleEvent) map[string]enums.PROCESS_STATUS {
	processStatusMap := make(map[string]enums.PROCESS_STATUS)
	for _, each := range events {
		if val, ok := processStatusMap[each.ProcessId]; ok {
			if each.Status == enums.FAILED {
				processStatusMap[each.ProcessId] = each.Status
			} else if val != enums.FAILED {
				if val == enums.PAUSED && each.Status == enums.ACTIVE {
					processStatusMap[each.ProcessId] = enums.ACTIVE
				} else if val == enums.COMPLETED && each.Status == enums.NON_INITIALIZED {
					processStatusMap[each.ProcessId] = enums.PAUSED
				} else if (val == enums.NON_INITIALIZED || val == enums.COMPLETED) && (each.Status != enums.NON_INITIALIZED && each.Status != enums.COMPLETED) {
					processStatusMap[each.ProcessId] = each.Status
				}
			}
		} else {
			processStatusMap[each.ProcessId] = each.Status
		}
	}
	return processStatusMap
}

func (p pipelineService) GetByProcessId(processId string) v1.Pipeline {
	events := p.processLifeCycleEventRepository.GetByProcessId(processId)
	if len(events) < 0 {
		return v1.Pipeline{}
	}
	var pipeline *v1.Pipeline
	if len(events) > 0 {
		pipeline = events[0].Pipeline
	}
	statusMap := make(map[string]enums.PROCESS_STATUS)
	for _, eachEvent := range events {
		key := eachEvent.Step + ":" + string(eachEvent.StepType)
		statusMap[key] = eachEvent.Status
	}
	for idx, _ := range pipeline.Steps {
		key := pipeline.Steps[idx].Name + ":" + string(pipeline.Steps[idx].Type)
		if status, ok := statusMap[key]; ok {
			pipeline.Steps[idx].Status = status
		}
	}
	return *pipeline
}

// NewPipelineService returns Pipeline type service
func NewPipelineService(processLifeCycleEventRepository repository.ProcessLifeCycleEventRepository) service.Pipeline {
	return &pipelineService{
		processLifeCycleEventRepository: processLifeCycleEventRepository,
	}
}

package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/klovercloud-ci-cd/event-bank/enums"
)

type pipelineService struct {
	processLifeCycleEventRepository repository.ProcessLifeCycleEventRepository
}

func (p pipelineService) GetByProcessId(processId string) v1.Pipeline {
	events := p.processLifeCycleEventRepository.GetByProcessId(processId)
	pipeline := events[0].Pipeline
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

